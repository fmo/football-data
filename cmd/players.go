package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var PlayersCmd = &cobra.Command{
	Use:   "players",
	Short: "Get players",
	Run: func(cmd *cobra.Command, args []string) {
		var players []maps.Player
		rapidPlayers := rapidapi.GetPlayers(0)
		maps.MapPlayers(rapidPlayers.Response, &players)

		for page := 1; page <= rapidPlayers.Paging.Total; page++ {
			result := rapidapi.GetPlayers(page)

			maps.MapPlayers(result.Response, &players)
		}

		for _, player := range players {
			fmt.Println(player.Name)
		}
	},
}
