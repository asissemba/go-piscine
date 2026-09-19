package main

import (
	"log"
	"net/http"

	"groupie-tracker-filters/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/artist", handler.ArtistHandler)
	mux.HandleFunc("/search", handler.FilterHandler)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
