package handler

import (
	"bytes"
	"html/template"

	"groupie-tracker-filters/internal/model"
)

func renderArtistCards(tpl *template.Template, artists []model.Artist) string {
	var buf bytes.Buffer

	if err := tpl.Execute(&buf, artists); err != nil {
		return ""
	}

	return buf.String()
}
