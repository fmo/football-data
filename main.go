package main

import (
	"fmt"
	"github.com/fmo/football-data/cmd"
	"github.com/fmo/football-data/cmd/games"
	"github.com/fmo/football-data/cmd/players"
	"github.com/fmo/football-data/cmd/teams"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	rootCmd := &cobra.Command{
		Use:   "football-data-app",
		Short: "Football Data CLI Application",
	}

	rootCmd.AddCommand(games.Cmd)
	rootCmd.AddCommand(cmd.CountriesCmd)
	rootCmd.AddCommand(cmd.LeaguesCmd)
	rootCmd.AddCommand(teams.Cmd)
	rootCmd.AddCommand(players.Cmd)
	rootCmd.AddCommand(cmd.StandingsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
