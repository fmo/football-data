package teams

import (
	"fmt"
	apifootballmap "github.com/fmo/football-data/internal/maps/apifootball"
	"github.com/fmo/football-data/internal/rapidapi/apifootball"
	"github.com/spf13/cobra"
)

var leagueId int
var teamId int

var Cmd = &cobra.Command{
	Use:   "teams",
	Short: "Get teams",
	Run: func(cmd *cobra.Command, args []string) {
		rapidTeams := apifootball.GetTeams(leagueId, teamId)
		teams := apifootballmap.MapTeams(rapidTeams)
		for _, team := range teams {
			fmt.Println(team.Id, team.Name)
		}
	},
}

func init() {
	Cmd.Flags().IntVarP(&leagueId, "leagueId", "l", 0, "League ID")
	Cmd.Flags().IntVarP(&teamId, "teamId", "t", 541, "Team ID")
}
