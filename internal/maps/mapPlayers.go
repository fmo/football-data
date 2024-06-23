package maps

import (
	"github.com/fmo/football-data/rapidapi"
	"strconv"
)

type Player struct {
	Team        string `json:"team"`
	TeamId      int    `json:"teamId"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Photo       string `json:"photo"`
	RapidApiID  string `json:"id"`
	Appearances int    `json:"appearances"`
	Position    string `json:"position"`
}

func MapPlayers(players []rapidapi.PlayerDetails, returnPlayer *[]Player) {
	for _, p := range players {
		player := Player{
			Name:        p.Player.Name,
			Firstname:   p.Player.Firstname,
			Lastname:    p.Player.Lastname,
			Age:         p.Player.Age,
			Nationality: p.Player.Nationality,
			Photo:       p.Player.Photo,
			RapidApiID:  strconv.Itoa(p.Player.ID),
			Team:        p.Statistics[0].Team.Name,
			TeamId:      p.Statistics[0].Team.ID,
			Appearances: p.Statistics[0].GameStats.Appearances,
			Position:    p.Statistics[0].GameStats.Position,
		}
		*returnPlayer = append(*returnPlayer, player)
	}
}
