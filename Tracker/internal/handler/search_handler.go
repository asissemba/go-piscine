package handler

import (
	"groupie-tracker-filters/internal/service"
	"net/http"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	result, _ := service.SearchArtists(query)

	w.Write([]byte("search ok"))
	_ = result
}
