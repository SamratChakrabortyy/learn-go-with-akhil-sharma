package datastore

import (
	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/models"
)

func GetAllPlayers() []models.Players {
	pSlice := []models.Players{}

	for _, pla := range players {
		pSlice = append(pSlice, *pla)
	}
	return pSlice
}
