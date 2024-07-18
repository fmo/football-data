package cmd

import (
	"fmt"
	"github.com/fmo/football-data/internal/rapidapi/apifootball"
	"github.com/spf13/cobra"
)

var CountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Get countries",
	Run: func(cmd *cobra.Command, args []string) {
		countries := apifootball.GetCountries()
		for _, country := range countries {
			fmt.Println(country.Name)
		}
	},
}
