package cmd

import (
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var leagueId int

var GamesCmd = &cobra.Command{
	Use:   "games",
	Short: "Get games",
	Run: func(cmd *cobra.Command, args []string) {
		requestParams := map[string]int{}
		requestParams["leagueId"] = leagueId
		requestParams["teamId"] = 541
		requestParams["season"] = 2023
		rapidApiGames := rapidapi.GetGames(requestParams)
		maps.MapGames(rapidApiGames)
		//publish(mappedGames, os.Getenv("KAFKA_TOPIC_FIXTURE"))
	},
}

func init() {
	GamesCmd.Flags().IntVarP(&leagueId, "leagueId", "l", 0, "League ID")
}
