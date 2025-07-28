package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("✨ Server humming softly, waiting for your truths.")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
