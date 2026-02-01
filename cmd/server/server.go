package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Serving at http://localhost:8080")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
}
