package transfermarkt

import (
	"fmt"
	transfermarktmap "github.com/fmo/football-data/internal/maps/transfermarkt"
	"github.com/fmo/football-data/internal/rapidapi/transfermarkt"
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
	Use:   "players-tm",
	Short: "Get players from Transfermarkt",
	Run: func(cmd *cobra.Command, args []string) {
		var players []*pb.Player

		log.Info("starting player command")

		r := transfermarkt.NewPlayersApi(log)
		rapidPlayers := r.GetPlayers()

		m := transfermarktmap.NewMapPlayers(log)
		m.MapPlayers(rapidPlayers, &players)

		fmt.Println(players)
	},
}
