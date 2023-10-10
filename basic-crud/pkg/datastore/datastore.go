package datastore

import "github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/models"

var players map[int]*models.Players = make(map[int]*models.Players)

var teams map[int]*models.Team = make(map[int]*models.Team)

func init() {
	initTeams()
	initPlayers()
}

func initTeams() {
	teams[1] = &models.Team{
		Id:   1,
		Name: "team 1",
		City: "city 1",
	}
	teams[2] = &models.Team{
		Id:   2,
		Name: "team 2",
		City: "city 2",
	}
}

func initPlayers() {
	players[1] = &models.Players{
		Id:      1,
		Name:    "pla 1",
		Country: "country 1",
		Team:    teams[1],
	}
	players[2] = &models.Players{
		Id:      2,
		Name:    "pla 2",
		Country: "country 2",
		Team:    teams[2],
	}
}
