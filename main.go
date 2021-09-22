package main

import (
	"fmt"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
}

func setupRoutes()  {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":80",nil)

}

func main()  {
	fmt.Println("Go File Upload Download Ade, test commit 3, tapi belum masuk")

}