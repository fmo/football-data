package maps

import (
	"github.com/fmo/football-data/rapidapi"
	pb "github.com/fmo/football-proto/golang/player"
	"github.com/sirupsen/logrus"
	"strconv"
)

type MapPlayersObj struct {
	logger *logrus.Logger
}

func NewMapPlayers(l *logrus.Logger) MapPlayersObj {
	return MapPlayersObj{
		logger: l,
	}
}

func (m MapPlayersObj) MapPlayers(players []rapidapi.PlayerDetails, returnPlayer *[]*pb.Player) {
	if len(players) == 0 {
		m.logger.Info("No players to map")
		return
	}

	m.logger.WithFields(logrus.Fields{
		"playerCountToMap": len(players),
		"sourceFile":       "mapPlayers",
		"function":         "MapPlayers",
	}).Info("mapping starting, number of players will be mapped")

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
