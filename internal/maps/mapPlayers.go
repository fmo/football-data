package maps

import (
	"github.com/fmo/football-data/rapidapi"
	pb "github.com/fmo/football-proto/golang/player"
	"strconv"
)

func MapPlayers(players []rapidapi.PlayerDetails, returnPlayer *[]*pb.Player) {
	for _, p := range players {
		player := &pb.Player{
			Name:        p.Player.Name,
			Firstname:   p.Player.Firstname,
			Lastname:    p.Player.Lastname,
			Age:         p.Player.Age,
			Nationality: p.Player.Nationality,
			Photo:       p.Player.Photo,
			RapidApiId:  strconv.Itoa(p.Player.ID),
			Team:        p.Statistics[0].Team.Name,
			TeamId:      p.Statistics[0].Team.ID,
			Appearances: p.Statistics[0].GameStats.Appearances,
			Position:    p.Statistics[0].GameStats.Position,
		}

		*returnPlayer = append(*returnPlayer, player)
	}
}
