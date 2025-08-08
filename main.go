package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT") // Railway provides this

	if port == "" {
		port = "8080" // fallback for local dev
	}

	log.Printf("✨ Server humming softly, waiting for your truths at http://localhost:%s", port)

	router := NewRouter()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
