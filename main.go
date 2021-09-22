package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	//1. parse input, multipart/form-data
	r.ParseMultipartForm (10 << 20)

	//2. rerive file from posted form-data
	file, handler, err := r.FormFile("MyFile")
	if err !=nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded File: %+v\n", handler.Filename)
	fmt.Println("File Size: %+v\n", handler.Size)
	fmt.Println("MIME Header: %+v\n", handler.Header)

	//3. write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images","upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err :=ioutil.ReadAll(file)
	if err !=nil{
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	//4. return wheter or not this has been successfull
	fmt.Println(w, "Successfully Uploaded File")

}

func setupRoutes()  {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe("localhost:8080",nil)

}

func main()  {
	fmt.Println("Upload Test Webpage")

}