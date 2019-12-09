package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const maxBytes = 2 * 1024 * 1024 // 2 MB
const uploadPath = "./tmp"

func randToken(len int) string {
	b := make([]byte, len)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.MaxBytesReader(w, req.Body, maxBytes)
		if err := req.ParseMultipartForm(maxBytes); err != nil {
			log.Printf("req.ParseMultipartForm: %v", err)
			return
		}

		file, _, err := req.FormFile("uploadFile")
		if err != nil {
			log.Printf("req.FormFile: %v", err)
			return
		}
		defer func() { _ = file.Close()}()

		f, _ := os.Create("poloxue.txt")
		defer func() {_ = f.Close()}()
		_, _ = io.Copy(f, file)

		fmt.Println(req.FormValue("words"))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tpl, err := template.ParseFiles("./form.html")
		if err != nil {
			log.Printf("template.New: %v", err)
			return
		}

		if err := tpl.Execute(w, nil); err != nil {
			log.Printf("tpl.Execute: %v", err)
			return
		}
	})
	http.HandleFunc("/upload", uploadHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files", http.StripPrefix("/files", fs))

	log.Println("Server started on localhost:8080, use /upload for uploading files and /files/{filename} for downloading files.")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
