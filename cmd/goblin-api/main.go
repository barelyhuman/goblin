package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/barelyhuman/goblin/build"
)

var shTemplates *template.Template
var serverURL string

// FIXME: Disabled storage and caching for initial version
// var storageClient *storage.Storage

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if path == "/" {
		path = "./static/index.html"
		http.ServeFile(rw, req, path)
		return
	}

	if strings.HasPrefix(path, "/static") {
		path = "./" + path
		_, err := os.Stat(path)
		if err != nil {
			rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
			rw.WriteHeader(http.StatusTeapot)
			fmt.Fprintln(rw, "Failed getting static file")
		}
		http.ServeFile(rw, req, path)
		return
	}

	if strings.HasPrefix(path, "/binary") {
		fetchBinary(rw, req)
		return
	}

	fetchInstallScript(rw, req)
}

func BlankReq(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	rw.Header().Set("Link", "rel=\"shortcut icon\" href=\"#\"")
}

func StartServer() {
	http.Handle("/favicon.ico", http.HandlerFunc(BlankReq))
	http.Handle("/", http.HandlerFunc(HandleRequest))

	port := envDefault("PORT", "3000")

	fmt.Println(">> Listening on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func envDefault(key string, def string) string {
	if s := os.Getenv(key); s == "" {
		return def
	} else {
		return s
	}
}

// TODO: cleanup code
// TODO: move everything into their own interface/structs
func main() {
	shTemplates = template.Must(template.ParseGlob("templates/*"))
	serverURL = envDefault("ORIGIN_URL", "http://localhost:3000")

	// FIXME: Disabled storage and caching for initial version
	// storageClient = &storage.Storage{}
	// storageClient.BucketName = os.Getenv("BUCKET_NAME")
	// err := storageClient.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	StartServer()
}

func normalizePackage(pkg string) string {
	// strip leading protocol
	pkg = strings.Replace(pkg, "https://", "", 1)

	// cleanup existing github.com
	pkg = strings.Replace(pkg, "github.com/", "", 1)

	// implicit github.com
	pkg = "github.com/" + pkg

	return pkg
}

func parsePackage(path string) (pkg, mod, version, bin string) {
	p := strings.Split(path, "@")
	version = "master"

	// pkg
	pkg = normalizePackage(p[0])

	// mod
	modp := strings.Split(pkg, "/")
	fmt.Println("modp", modp)
	if len(modp) >= 3 {
		mod = strings.Join(modp[:3], "/")
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
		Version:         version,
	})
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

	bin := &build.Binary{
		Path:    pkg,
		Version: version,
		OS:      goos,
		Arch:    arch,
		Name:    name,
		Module:  mod,
	}

	// TODO: check the storage for existing binary for the module
	// and return from the storage instead

	immutable(rw)

	// FIXME: Disabled storage and caching for initial version
	// var buf bytes.Buffer
	// err := bin.WriteBuild(io.MultiWriter(rw, &buf))

	err := bin.WriteBuild(io.MultiWriter(rw))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err.Error())
		return
	}

	err = bin.Cleanup()
	if err != nil {
		log.Println("cleaning binary build", err)
	}

	// FIXME: Disabled storage and caching for initial version
	// err = storageClient.Upload(bin.Module, bin.Dest)
	// if err != nil {
	// 	fmt.Fprint(rw, err.Error())
	// 	return
	// }

	// url, err := storageClient.GetSignedURL(bin.Module, bin.Name)
	// if err != nil {
	// 	fmt.Fprint(rw, err.Error())
	// 	return
	// }
}
