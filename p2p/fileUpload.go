package main

import (
	"fmt"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Uploading File")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
}

func main() {
    fmt.Println("Server Run on Port 8080")
    setupRoutes()
}
