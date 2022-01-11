package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/barelyhuman/goblin/build"
	"github.com/barelyhuman/goblin/storage"
)

var storageClient *storage.Storage

func HandleRequest(rw http.ResponseWriter, req *http.Request) {
	// Hard coded for dev

	bin := &build.Binary{
		Path:    "github.com/barelyhuman/commitlog",
		Version: "v0.0.6",
		OS:      "linux",
		Arch:    "amd64",
		Module:  "github.com/barelyhuman/commitlog",
	}

	err := bin.WriteBuild()
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	err = storageClient.Upload(bin.Module, bin.Dest)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}

	url, err := storageClient.GetSignedURL(bin.Module)
	if err != nil {
		fmt.Fprint(rw, err.Error())
		return
	}
	fmt.Fprintf(rw, "url:%v", url)
}

func StartServer() {
	http.Handle("/", http.HandlerFunc(HandleRequest))
	port := "3000"
	fmt.Println(">> Listening on " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	storageClient = &storage.Storage{}
	storageClient.BucketName = os.Getenv("BUCKET_NAME")
	err := storageClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	StartServer()
}
