package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var TeamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Get teams",
	Run: func(cmd *cobra.Command, args []string) {
		rapidTeams := rapidapi.GetTeams(2)
		teams := maps.MapTeams(rapidTeams)
		for _, team := range teams {
			fmt.Println(team.Name)
		}
	},
}
