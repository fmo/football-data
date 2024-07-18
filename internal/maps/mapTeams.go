package maps

import (
	"github.com/fmo/football-data/internal/rapidapi"
	"strconv"
)

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func MapTeams(rapidTeams []rapidapi.TeamsSpecs) []Team {
	var teams []Team
	for _, t := range rapidTeams {
		teams = append(teams, Team{
			Id:   strconv.Itoa(t.Team.ID),
			Name: t.Team.Name,
			Logo: t.Team.Logo,
		})
	}

	return teams
}
