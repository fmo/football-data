package transfermarkt

import (
	"github.com/fmo/football-data/internal/rapidapi/transfermarkt"
	pb "github.com/fmo/football-proto/golang/player"
	"github.com/sirupsen/logrus"
)

type MapPlayersObj struct {
	logger *logrus.Logger
}

func NewMapPlayers(l *logrus.Logger) MapPlayersObj {
	return MapPlayersObj{
		logger: l,
	}
}

func (m MapPlayersObj) MapPlayers(players []transfermarkt.Player, returnPlayer *[]*pb.Player, teamId int) {
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
			Name:                p.Name,
			TransfermarktId:     p.ID,
			ShirtNumber:         p.ShirtNumber,
			MarketValue:         int32(p.MarketValue.Value),
			MarketValueCurrency: p.MarketValue.Currency,
			TeamId:              int32(teamId),
		}

		*returnPlayer = append(*returnPlayer, player)
	}
}
