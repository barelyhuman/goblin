package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/barelyhuman/go/env"
	"github.com/barelyhuman/goblin/build"
	"github.com/barelyhuman/goblin/resolver"
	"github.com/barelyhuman/goblin/storage"
	"github.com/joho/godotenv"
)

var shTemplates *template.Template
var serverURL string
var storageClient storage.Storage

type ErrorJSON struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (ej ErrorJSON) toJSONString() (string, error) {
	marshaled, err := json.Marshal(ej)
	if err != nil {
		return "", err
	}
	return string(marshaled), nil
}

type VersionJSON struct {
	Success         bool   `json:"success"`
	Package         string `json:"package"`
	Binary          string `json:"binary"`
	OriginalVersion string `json:"originalVersion"`
	Version         string `json:"version"`
}

func (ej VersionJSON) toJSONString() (string, error) {
	marshaled, err := json.Marshal(ej)
	if err != nil {
		return "", err
	}
	return string(marshaled), nil
}

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		path = "./static/index.html"
		http.ServeFile(rw, req, path)
		return
	}

	file := filepath.Join("static", path)
	info, err := os.Stat(file)
	if err == nil && info.Mode().IsRegular() {
		http.ServeFile(rw, req, file)
		return
	}

	if strings.HasPrefix(path, "/version") {
		log.Println("Resolving version")
		resolveVersionJSON(rw, req)
		return
	}

	if strings.HasPrefix(path, "/binary") {
		log.Print("handle binary")
		fetchBinary(rw, req)
		return
	}

	fetchInstallScript(rw, req)
}

func BlankReq(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	rw.Header().Set("Link", "rel=\"shortcut icon\" href=\"#\"")
}

func StartServer(port string) {
	http.Handle("/favicon.ico", http.HandlerFunc(BlankReq))
	http.Handle("/", http.HandlerFunc(HandleRequest))

	fmt.Println(">> Listening on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: cleanup code
// TODO: move everything into their own interface/structs
func main() {

	envFile := flag.String("env", ".env", "path to read the env config from")
	portFlag := env.Get("PORT", "3000")

	flag.Parse()

	if _, err := os.Stat(*envFile); !errors.Is(err, os.ErrNotExist) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	shTemplates = template.Must(template.ParseGlob("templates/*"))
	serverURL = env.Get("ORIGIN_URL", "http://localhost:"+portFlag)

	if isStorageEnabled() {
		storageClient = storage.NewAWSStorage(env.Get("STORAGE_BUCKET", "goblin-cache"))
		err := storageClient.Connect()
		if err != nil {
			log.Fatal(err)
		}
	}

	clearStorageBackgroundJob()
	StartServer(portFlag)
}

func clearStorageBackgroundJob() {
	cacheHoldEnv := env.Get("CLEAR_CACHE_TIME", "")
	if len(cacheHoldEnv) == 0 {
		return
	}

	cacheHoldDuration, _ := time.ParseDuration(cacheHoldEnv)

	cleaner := func(storageClient storage.Storage) {
		log.Println("Cleaning Cached Storage Object")
		objects := storageClient.ListObjects()
		for _, obj := range objects {
			objExpiry := obj.LastModified.Add(cacheHoldDuration)
			if time.Now().Equal(objExpiry) || time.Now().After(objExpiry) {
				storageClient.RemoveObject(obj.Key)
			}
		}
	}

	ticker := time.NewTicker(cacheHoldDuration)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				cleaner(storageClient)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func isStorageEnabled() bool {
	useStorageEnv := env.Get("STORAGE_ENABLED", "false")
	useStorage := false
	if useStorageEnv == "true" {
		useStorage = true
	}
	return useStorage
}

func normalizePackage(pkg string) string {
	// strip leading protocol
	pkg = strings.Replace(pkg, "https://", "", 1)
	return pkg
}

func parsePackage(path string) (pkg, mod, version, bin string) {
	p := strings.Split(path, "@")
	version = ""

	// pkg
	pkg = normalizePackage(p[0])

	// mod
	modp := strings.Split(pkg, "/")
	if len(modp) >= 3 {
		mod = strings.Join(modp[:3], "/")
	} else {
		mod = pkg
	}

	// version after @
	if len(p) > 1 {
		version = p[1]
	}

	// binary name from pkg
	p = strings.Split(pkg, "/")
	bin = p[len(p)-1]
	return
}

// immutable sets immutability header fields.
func immutable(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Cache-Control", "max-age=31536000, immutable")
}

func render(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "application/x-sh")
	w.Header().Set("Cache-Control", "no-store")
	shTemplates.ExecuteTemplate(w, name, data)
}

func fetchInstallScript(rw http.ResponseWriter, req *http.Request) {
	pkg := strings.TrimPrefix(req.URL.Path, "/")
	pkg, _, version, name := parsePackage(pkg)

	v := &resolver.Resolver{
		Pkg: pkg,
	}

	v.ParseVersion(version)
	resolvedVersion, err := v.ResolveVersion()
	if err != nil || len(resolvedVersion) == 0 {
		render(rw, "error.sh", ("Failed to resolve version:" + version))
		return
	}

	// == mark default to latest version when nothing is provided ==
	// this has be separated and put here since `latest` might actually
	// be a tag provided to the package
	// and could be then used, so using the branch name
	// makes no sense when working with go proxy instead of
	// github for example
	if len(version) == 0 {
		version = "latest"
	}

	render(rw, "install.sh", struct {
		URL             string
		Package         string
		Binary          string
		OriginalVersion string
		Version         string
	}{
		URL:             serverURL,
		Package:         pkg,
		Binary:          name,
		OriginalVersion: version,
		Version:         resolvedVersion,
	})
}

func resolveVersionJSON(rw http.ResponseWriter, req *http.Request) {
	// only reply in JSON
	rw.Header().Set("Content-Type", "application/json")

	pkg := strings.TrimPrefix(req.URL.Path, "/version")
	pkg, _, version, name := parsePackage(pkg)
	v := &resolver.Resolver{
		Pkg: pkg,
	}
	v.ParseVersion(version)
	resolvedVersion, err := v.ResolveVersion()
	if err != nil || len(resolvedVersion) == 0 {
		errorJson, _ := ErrorJSON{Success: false, Message: "Failed to resolve version:" + version}.toJSONString()
		rw.Write([]byte(errorJson))
		return
	}

	responseJson, _ := VersionJSON{
		Success:         true,
		Package:         pkg,
		Binary:          name,
		OriginalVersion: version,
		Version:         resolvedVersion,
	}.toJSONString()

	rw.Write([]byte(responseJson))
	return

}

func fetchBinary(rw http.ResponseWriter, req *http.Request) {
	pkg := strings.TrimPrefix(req.URL.Path, "/binary/")

	pkg, mod, _, name := parsePackage(pkg)

	goos := req.URL.Query().Get("os")
	if goos == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "`os` is a required parameter")
		return
	}

	arch := req.URL.Query().Get("arch")
	if arch == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "`arch` is a required parameter")
		return
	}

	version := req.URL.Query().Get("version")
	if version == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "`version` is a required parameter")
		return
	}

	binName := req.URL.Query().Get("out")
	if binName == "" {
		binName = name
	}

	cmdPath := req.URL.Query().Get("cmd")

	bin := &build.Binary{
		Path:    pkg,
		Version: version,
		OS:      goos,
		CmdPath: cmdPath,
		Arch:    arch,
		Name:    binName,
		Module:  mod,
	}

	immutable(rw)

	artifactName := constructArtifactName(bin)

	if isStorageEnabled() && storageClient.HasObject(artifactName) {
		url, _ := storageClient.GetSignedURL(artifactName)
		log.Println("From cache")
		http.Redirect(rw, req, url, http.StatusSeeOther)
		return
	}

	var buf bytes.Buffer
	err := bin.WriteBuild(io.MultiWriter(rw, &buf))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err.Error())
		return
	}

	if isStorageEnabled() {
		err = storageClient.Upload(
			artifactName,
			buf,
		)

		if err != nil {
			log.Println("Failed to upload", err)
		}
	}

	err = bin.Cleanup()
	if err != nil {
		log.Println("cleaning binary build", err)
	}
}

func constructArtifactName(bin *build.Binary) string {
	var artifactName strings.Builder
	artifactName.Write([]byte(bin.Name))
	artifactName.Write([]byte("-"))
	artifactName.Write([]byte(bin.Version))
	artifactName.Write([]byte("-"))
	artifactName.Write([]byte(bin.OS))
	artifactName.Write([]byte("-"))
	artifactName.Write([]byte(bin.Arch))
	return artifactName.String()
}
