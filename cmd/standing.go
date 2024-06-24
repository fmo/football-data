package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var StandingsCmd = &cobra.Command{
	Use:   "standings",
	Short: "Get standings",
	Run: func(cmd *cobra.Command, args []string) {
		rapidStanding := rapidapi.GetStanding(2)
		standing := maps.MapStanding(rapidStanding)
		for _, s := range standing {
			fmt.Println(s.Name)
		}
	},
}
