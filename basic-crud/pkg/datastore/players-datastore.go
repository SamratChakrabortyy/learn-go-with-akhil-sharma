package datastore

import (
	"github.com/SamratChakrabortyy/learn-go-with-akhil-sharma/basic-crud/pkg/models"
)

func GetAllPlayers() []models.Player {
	pSlice := []models.Player{}

	for _, pla := range players {
		pSlice = append(pSlice, *pla)
	}
	return pSlice
}

func GetPlayerById(id int) models.Player {
	return *players[id]
}

func AddPlayer(p models.Player) {
	count := len(players)
	p.Id = count + 1
	if _, existingTeam := teams[p.Team.Id]; !existingTeam {
		count = len(teams)
		team := *p.Team
		team.Id = count + 1
		teams[team.Id] = &team
	}
	p.Team = teams[p.Team.Id]
	players[p.Id] = &p
}
