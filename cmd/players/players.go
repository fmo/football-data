package players

import (
	"github.com/fmo/football-data/internal/kafka"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	pb "github.com/fmo/football-proto/golang/player"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var teamId int
var season int

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func init() {
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 541, "Team Id")
	Cmd.Flags().IntVarP(&season, "season", "s", 2023, "Season")

	// Log as JSON instead of the default ASCII formatter.
	log.Formatter = &logrus.JSONFormatter{}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	// Only log the warning severity or above.
	log.Level = logrus.DebugLevel
}

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Get players",
	Run: func(cmd *cobra.Command, args []string) {
		var players *[]*pb.Player
		rapidPlayers := rapidapi.GetPlayers(0, season, teamId)
		maps.MapPlayers(rapidPlayers.Response, players)

		for page := 1; page <= rapidPlayers.Paging.Total; page++ {
			result := rapidapi.GetPlayers(page, season, teamId)

			maps.MapPlayers(result.Response, players)
		}

		kafka.Publish(log, players, os.Getenv("KAFKA_TOPIC_PLAYERS"))
	},
}
