package players

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var teamId int
var season int

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Get players",
	Run: func(cmd *cobra.Command, args []string) {
		var players []maps.Player
		rapidPlayers := rapidapi.GetPlayers(0, season, teamId)
		maps.MapPlayers(rapidPlayers.Response, &players)

		for page := 1; page <= rapidPlayers.Paging.Total; page++ {
			result := rapidapi.GetPlayers(page, season, teamId)

			maps.MapPlayers(result.Response, &players)
		}

		for _, player := range players {
			fmt.Println(player.Name)
		}
	},
}

func init() {
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 541, "Team Id")
	Cmd.Flags().IntVarP(&season, "season", "s", 2023, "Season")
}
