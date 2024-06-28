package maps

import (
	"github.com/fmo/football-data/rapidapi"
)

type responseTeam struct {
	LeagueId     int    `json:"leagueId"`
	LeagueName   string `json:"leagueName"`
	TeamId       int    `json:"teamId"`
	Name         string `json:"name"`
	Rank         int    `json:"rank"`
	Group        string `json:"group"`
	Points       int    `json:"points"`
	Played       int    `json:"played"`
	Win          int    `json:"win"`
	Lose         int    `json:"lose"`
	GoalsFor     int    `json:"goalsFor"`
	GoalsAgainst int    `json:"goalsAgainst"`
}

func MapStanding(rapidStanding rapidapi.StandingsResponse) []responseTeam {
	var responseTeams []responseTeam
	for _, league := range rapidStanding.Response {
		for _, standing := range league.League.Standings {
			for _, team := range standing {
				responseTeams = append(responseTeams, responseTeam{
					LeagueId:     league.League.ID,
					LeagueName:   league.League.Name,
					Name:         team.Team.Name,
					Rank:         team.Rank,
					Points:       team.Points,
					Played:       team.All.Played,
					Win:          team.All.Win,
					Lose:         team.All.Lose,
					GoalsFor:     team.All.Goals.For,
					GoalsAgainst: team.All.Goals.Against,
					Group:        team.Group,
					TeamId:       team.Team.ID,
				})
			}
		}
	}

	return responseTeams
}
