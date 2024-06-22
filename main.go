package main

import (
	"fmt"
	"github.com/fmo/football-data/rapidapi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	requestParams := map[string]int{}

	actionId := getActionId()
	leagueId := getLeagueId()
	teamId := getTeamId()
	season := getSeason()

	log.WithFields(
		log.Fields{
			"actionId": actionId,
			"leagueId": leagueId,
			"teamId":   teamId,
			"season":   season,
		},
	).Info("Entered Inputs")

	requestParams["leagueId"] = leagueId
	requestParams["season"] = season
	if teamId != 0 {
		requestParams["teamId"] = teamId
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	switch actionId {
	case 0:
		rapidApiGames := rapidapi.GetGames(requestParams)
		mapGames(rapidApiGames, teamId)
		//publish(mappedGames, os.Getenv("KAFKA_TOPIC_FIXTURE"))
	case 1:
		countries := rapidapi.GetCountries()
		for _, country := range countries {
			fmt.Println(country.Name)
		}
	case 2:
		leagues := rapidapi.GetLeagues()
		for _, league := range leagues {
			fmt.Println(league.League.Name)
		}
	case 3:
		var players []Player
		rapidPlayers := rapidapi.GetPlayers(0)
		mapPlayers(rapidPlayers.Response, &players)

		for page := 1; page <= rapidPlayers.Paging.Total; page++ {
			result := rapidapi.GetPlayers(page)

			mapPlayers(result.Response, &players)
		}

		for _, player := range players {
			fmt.Println(player.Name)
		}
	case 4:
		rapidTeams := rapidapi.GetTeams(leagueId)
		teams := MapTeams(rapidTeams)
		for _, team := range teams {
			fmt.Println(team.Name)
		}
	case 5:
		rapidStanding := rapidapi.GetStanding(leagueId)
		standing := MapStanding(rapidStanding)
		for _, s := range standing {
			fmt.Println(s.Name)
		}
	}
}
