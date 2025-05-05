package main

import (
	"log"
	"net/http"

	"github.com/maxzhirnov/utm-builder/internal/config"
	"github.com/maxzhirnov/utm-builder/internal/handlers"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Serve static files from the assets directory
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Application routes
	http.HandleFunc("/", handlers.FormHandler)
	http.HandleFunc("/generate", handlers.GenerateHandler)
	http.HandleFunc("/options", handlers.GetOptionsHandler)

	log.Println("Server starting at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
