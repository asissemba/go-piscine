package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie-tracker-filters/internal/model"
)

var cachedArtists []model.Artist
var cachedRelations []model.Relation

func GetArtists() ([]model.Artist, error) {

	if cachedArtists != nil {
		return cachedArtists, nil
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&cachedArtists)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return cachedArtists, nil
}

func GetRelations() ([]model.Relation, error) {

	if cachedRelations != nil {
		return cachedRelations, nil
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// ← ВСТАВИТЬ СЮДА
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var response model.RelationsResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	cachedRelations = response.Index

	return cachedRelations, nil
}
