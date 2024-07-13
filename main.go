package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/secret", getSecret)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
