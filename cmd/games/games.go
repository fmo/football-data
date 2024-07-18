package games

import (
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/internal/rapidapi/apifootball"
	"github.com/spf13/cobra"
)

var leagueId int
var teamId int
var season int

var Cmd = &cobra.Command{
	Use:   "games",
	Short: "Get games",
	Run: func(cmd *cobra.Command, args []string) {
		requestParams := map[string]int{}
		requestParams["leagueId"] = leagueId
		requestParams["teamId"] = teamId
		requestParams["season"] = season
		rapidApiGames := apifootball.GetGames(requestParams)
		maps.MapGames(rapidApiGames)
		//publish(mappedGames, os.Getenv("KAFKA_TOPIC_FIXTURE"))
	},
}

func init() {
	Cmd.Flags().IntVarP(&leagueId, "leagueId", "l", 0, "League ID")
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 541, "Team ID")
	Cmd.Flags().IntVarP(&season, "season", "s", 2023, "Season")
}
