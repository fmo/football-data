package standings

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/internal/rapidapi/apifootball"
	"github.com/spf13/cobra"
)

var leagueId int
var season int

var Cmd = &cobra.Command{
	Use:   "standings",
	Short: "Get standings",
	Run: func(cmd *cobra.Command, args []string) {
		rapidStanding := apifootball.GetStanding(leagueId, season)
		standing := maps.MapStanding(rapidStanding)
		for _, s := range standing {
			fmt.Println(s.Name)
		}
	},
}

func init() {
	Cmd.Flags().IntVarP(&leagueId, "leagueId", "l", 2, "League Id")
	Cmd.Flags().IntVarP(&season, "season", "s", 2023, "Season")
}
