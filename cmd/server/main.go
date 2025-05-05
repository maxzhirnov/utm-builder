package main

import (
	"log"
	"net/http"

	"github.com/maxzhirnov/utm-builder/internal/handlers"
)

func main() {
	// Serve static files from the assets directory
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Handle application routes
	http.HandleFunc("/", handlers.FormHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
