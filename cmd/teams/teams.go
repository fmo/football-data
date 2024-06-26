package teams

import (
	"fmt"
	"github.com/fmo/football-data/internal/maps"
	"github.com/fmo/football-data/rapidapi"
	"github.com/spf13/cobra"
)

var leagueId int
var teamId int

var Cmd = &cobra.Command{
	Use:   "teams",
	Short: "Get teams",
	Run: func(cmd *cobra.Command, args []string) {
		rapidTeams := rapidapi.GetTeams(leagueId, teamId)
		teams := maps.MapTeams(rapidTeams)
		for _, team := range teams {
			fmt.Println(team.Id, team.Name)
		}
	},
}

func init() {
	Cmd.Flags().IntVarP(&leagueId, "leagueId", "l", 0, "League ID")
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 541, "Team ID")
}
