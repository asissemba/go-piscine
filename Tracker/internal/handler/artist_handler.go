package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"groupie-tracker-filters/internal/service"
)

var artistTemplate = template.Must(
	template.ParseFiles(templatePath("artist.html")),
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist id", http.StatusBadRequest)
		return
	}

	artist, err := service.GetArtistByID(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if artist == nil {
		http.NotFound(w, r)
		return
	}

	relation, err := service.GetRelationByArtistID(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Artist   any
		Relation any
	}{
		Artist:   artist,
		Relation: relation,
	}

	if err := artistTemplate.Execute(w, data); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
}
