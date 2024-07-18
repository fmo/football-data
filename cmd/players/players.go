package players

import (
	"github.com/fmo/football-data/internal/kafka"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/internal/rapidapi"
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

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	// Only log the debug severity or above.
	log.Level = logrus.DebugLevel
}

var Cmd = &cobra.Command{
	Use:   "players",
	Short: "Get players",
	Run: func(cmd *cobra.Command, args []string) {
		var players []*pb.Player

		log.Info("starting player command")

		r := rapidapi.NewRapidApi(log)
		rapidPlayers := r.GetPlayers(0, season, teamId)

		m := maps.NewMapPlayers(log)
		m.MapPlayers(rapidPlayers.Response, &players)

		log.WithFields(logrus.Fields{
			"numberOfMappedPlayers": len(players),
			"sourceFile":            "cmd/players/players.go",
			"function":              "command",
		}).Info("mapping completed for the first page")

		for page := 2; page <= rapidPlayers.Paging.Total; page++ {
			result := r.GetPlayers(page, season, teamId)

			m.MapPlayers(result.Response, &players)

			log.WithFields(logrus.Fields{
				"numberOfMappedPlayers": len(players),
				"sourceFile":            "cmd/players/players.go",
				"function":              "command",
			}).Infof("mapping completed for the page %d", page)
		}

		publisher := kafka.NewPublisher(log, os.Getenv("KAFKA_TOPIC_PLAYERS"))
		publisher.Publish(players)
	},
}
