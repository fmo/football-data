package cmd

import (
	"fmt"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var CountriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Get countries",
	Run: func(cmd *cobra.Command, args []string) {
		countries := rapidapi.GetCountries()
		for _, country := range countries {
			fmt.Println(country.Name)
		}
	},
}
