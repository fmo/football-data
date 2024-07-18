package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/rapidapi"
	"github.com/spf13/cobra"
)

var LeaguesCmd = &cobra.Command{
	Use:   "leagues",
	Short: "Get leagues",
	Run: func(cmd *cobra.Command, args []string) {
		leagues := rapidapi.GetLeagues()
		for _, league := range leagues {
			fmt.Println(league.League.Name)
		}
	},
}
