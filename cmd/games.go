package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/leagues"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var GamesCmd = &cobra.Command{
	Use:   "games",
	Short: "Get games",
	Run: func(cmd *cobra.Command, args []string) {
		requestParams := map[string]int{}
		requestParams["leagueId"] = getLeagueId()
		requestParams["teamId"] = 541
		requestParams["season"] = 2023
		rapidApiGames := rapidapi.GetGames(requestParams)
		maps.MapGames(rapidApiGames)
		//publish(mappedGames, os.Getenv("KAFKA_TOPIC_FIXTURE"))
	},
}

func getLeagueId() int {
	var leagueId int
	fmt.Println("Select a league")
	fmt.Println("----------------")
	for i, l := range leagues.GetLeagues() {
		fmt.Printf("%s - %d\n", i, l)
	}

	_, err := fmt.Scan(&leagueId)
	if err != nil {
		fmt.Println("Invalid input")
	}

	return leagueId
}
