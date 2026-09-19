package service

import (
	"strconv"
	"strings"

	"groupie-tracker-filters/internal/model"
)

func ApplyFilters(
	artists []model.Artist,
	relations []model.Relation,
	f model.FilterParams,
) []model.Artist {

	relationMap := make(map[int]model.Relation)

	for _, r := range relations {
		relationMap[r.ID] = r
	}

	result := make([]model.Artist, 0, len(artists))

	for _, a := range artists {

		if f.Search != "" &&
			!strings.Contains(
				strings.ToLower(a.Name),
				strings.ToLower(f.Search),
			) {
			continue
		}

		// 1. Creation date range
		if f.CreationFrom != 0 && a.CreationDate < f.CreationFrom {
			continue
		}
		if f.CreationTo != 0 && a.CreationDate > f.CreationTo {
			continue
		}

		// 2. First album year
		if f.AlbumFrom != 0 || f.AlbumTo != 0 {

			if len(a.FirstAlbum) >= 4 {

				year, err := strconv.Atoi(a.FirstAlbum[len(a.FirstAlbum)-4:])
				if err == nil {

					if f.AlbumFrom != 0 && year < f.AlbumFrom {
						continue
					}

					if f.AlbumTo != 0 && year > f.AlbumTo {
						continue
					}
				}
			}
		}

		// 3. Members checkbox
		if len(f.Members) > 0 {

			ok := false

			for _, m := range f.Members {
				if len(a.Members) == m {
					ok = true
					break
				}
			}

			if !ok {
				continue
			}
		}

		// 4. Location filter
		if len(f.Locations) > 0 {
			found := false

			relation, ok := relationMap[a.ID]
			if !ok {
				continue
			}

			for _, loc := range f.Locations {
				loc = strings.ToLower(loc)

				for city := range relation.DatesLocations {
					if strings.Contains(strings.ToLower(city), loc) {
						found = true
						break
					}
				}

				if found {
					break
				}
			}

			if !found {
				continue
			}
		}

		// 5. Number of Locations
		if f.LocationCount != 0 {
			relation, ok := relationMap[a.ID]
			if !ok {
				continue
			}

			if len(relation.DatesLocations) != f.LocationCount {
				continue
			}
		}

		result = append(result, a)
	}

	return result
}
