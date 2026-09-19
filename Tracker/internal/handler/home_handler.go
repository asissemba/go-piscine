package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker-filters/internal/model"
	"groupie-tracker-filters/internal/service"
)

var tmpl = template.Must(
	template.New("home.html").Funcs(template.FuncMap{
		"containsInt": func(slice []int, value int) bool {
			for _, v := range slice {
				if v == value {
					return true
				}
			}
			return false
		},
	}).ParseFiles(templatePath("home.html")),
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := service.GetArtists()
	if err != nil {
		http.Error(w, "Failed to load artists", http.StatusInternalServerError)
		return
	}

	relations, err := service.GetRelations()
	if err != nil {
		http.Error(w, "Failed to load relations", http.StatusInternalServerError)
		return
	}

	var f model.FilterParams

	f.Search = r.URL.Query().Get("q")

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	albumFrom := r.URL.Query().Get("album_from")
	albumTo := r.URL.Query().Get("album_to")

	location := strings.TrimSpace(r.URL.Query().Get("location"))
	f.Location = location

	members := r.URL.Query()["members"]

	if from != "" {
		f.CreationFrom, _ = strconv.Atoi(from)
	}

	if to != "" {
		f.CreationTo, _ = strconv.Atoi(to)
	}

	if albumFrom != "" {
		f.AlbumFrom, _ = strconv.Atoi(albumFrom)
	}

	if albumTo != "" {
		f.AlbumTo, _ = strconv.Atoi(albumTo)
	}

	concertCount := r.URL.Query().Get("concert_count")

	if concertCount != "" {
		f.ConcertCount, _ = strconv.Atoi(concertCount)
	}

	if len(members) > 0 {
		for _, p := range members {
			m, err := strconv.Atoi(strings.TrimSpace(p))
			if err == nil {
				f.Members = append(f.Members, m)
			}
		}
	}

	if location != "" {
		f.Locations = append(f.Locations, location)
	}

	filtered := service.ApplyFilters(artists, relations, f)

	data := model.HomePageData{
		Artists: filtered,
		Filters: f,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
