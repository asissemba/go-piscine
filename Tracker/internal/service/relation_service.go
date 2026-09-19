package service

import (
	"groupie-tracker-filters/internal/client"
	"groupie-tracker-filters/internal/model"
)

func GetRelations() ([]model.Relation, error) {
	return client.GetRelations()
}
