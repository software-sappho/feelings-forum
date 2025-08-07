package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("✨ Server humming softly, waiting for your truths at http://localhost:8080")

	router := NewRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
