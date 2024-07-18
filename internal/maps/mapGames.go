package maps

import (
	"github.com/fmo/football-data/internal/rapidapi/apifootball"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type Game struct {
	FixtureId    int    `json:"fixtureId"`
	Home         string `json:"home"`
	HomeTeamId   int    `json:"homeTeamId"`
	Away         string `json:"away"`
	AwayTeamId   int    `json:"awayTeamId"`
	Status       string `json:"status"`
	Score        string `json:"score"`
	Round        string `json:"round"`
	Timestamp    int    `json:"timestamp"`
	GameDate     string `json:"date"`
	LeagueId     int    `json:"leagueId"`
	LeagueName   string `json:"leagueName"`
	HomeTeamLogo string `json:"homeTeamLogo"`
	AwayTeamLogo string `json:"awayTeamLogo"`
}

func MapGames(fixtures []apifootball.FixtureDetails) []Game {
	var games []Game

	for _, fixture := range fixtures {
		game := Game{
			Home:         fixture.Teams.Home.Name,
			HomeTeamId:   fixture.Teams.Home.ID,
			Away:         fixture.Teams.Away.Name,
			AwayTeamId:   fixture.Teams.Away.ID,
			Status:       fixture.Fixture.Status.Short,
			Score:        strconv.Itoa(fixture.Score.FullTime.Home) + "-" + strconv.Itoa(fixture.Score.FullTime.Away),
			Round:        fixture.League.Round,
			FixtureId:    fixture.Fixture.Id,
			Timestamp:    fixture.Fixture.Timestamp,
			GameDate:     fixture.Fixture.GameDate,
			LeagueId:     fixture.League.ID,
			LeagueName:   fixture.League.Name,
			HomeTeamLogo: fixture.Teams.Home.Logo,
			AwayTeamLogo: fixture.Teams.Away.Logo,
		}

		games = append(games, game)

		log.WithFields(log.Fields{
			"homeTeam":   game.Home,
			"awayTeam":   game.Away,
			"gameDate":   game.GameDate,
			"round":      game.Round,
			"leagueName": game.LeagueName,
			"score":      game.Score,
		}).Info("Game fetched!")
	}

	return games
}
