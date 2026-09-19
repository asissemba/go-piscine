package service

import (
	"strconv"
	"strings"

	"groupie-tracker-filters/internal/client"
	"groupie-tracker-filters/internal/model"
)

func GetArtists() ([]model.Artist, error) {
	return client.GetArtists()
}

func SearchArtists(query string) ([]model.Artist, error) {

	artists, err := client.GetArtists()
	if err != nil {
		return nil, err
	}

	var result []model.Artist

	query = strings.ToLower(query)

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			result = append(result, artist)
		}
	}

	return result, nil
}

func GetArtistByID(id int) (*model.Artist, error) {
	artists, err := client.GetArtists()
	if err != nil {
		return nil, err
	}

	for _, a := range artists {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, nil
}

func GetRelationByArtistID(id int) (*model.Relation, error) {

    relations, err := client.GetRelations()
    if err != nil {
        return nil, err
    }

    for _, r := range relations {
        if r.ID == id {
            return &r, nil
        }
    }

    return nil, nil
}

func FilterArtists(
	artists []model.Artist,
	search string,
	year string,
	members string,
) []model.Artist {

	var result []model.Artist

	for _, a := range artists {

		if search != "" &&
			!strings.Contains(
				strings.ToLower(a.Name),
				strings.ToLower(search),
			) {
			continue
		}

		if year != "" {
			y, err := strconv.Atoi(year)
			if err == nil && a.CreationDate != y {
				continue
			}
		}

		if members != "" {
			m, err := strconv.Atoi(members)
			if err == nil && len(a.Members) != m {
				continue
			}
		}

		result = append(result, a)
	}

	return result
}
