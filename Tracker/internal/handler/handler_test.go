package handler

import (
	"html/template"
	"path/filepath"
	"strings"
	"testing"

	"groupie-tracker-filters/internal/model"
)

func TestRenderArtistCards(t *testing.T) {
	tplPath := filepath.Join("..", "..", "templates", "cards.html")
	tpl := template.Must(template.ParseFiles(tplPath))

	artists := []model.Artist{{
		ID:           1,
		Name:         "Test Artist",
		Image:        "/img.png",
		CreationDate: 1999,
		Members:      []string{"A", "B"},
	}}

	html := renderArtistCards(tpl, artists)

	if !strings.Contains(html, "Test Artist") {
		t.Fatalf("expected rendered HTML to contain artist name, got %q", html)
	}

	if !strings.Contains(html, "/artist?id=1") {
		t.Fatalf("expected rendered HTML to contain artist link, got %q", html)
	}
}
