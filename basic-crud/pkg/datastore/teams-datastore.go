package datastore

import "github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/models"

func GetAllTeams() []models.Team {
	tSlice := []models.Team{}
	for _, team := range teams {
		tSlice = append(tSlice, *team)
	}
	return tSlice
}
