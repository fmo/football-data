package rapidapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type StandingsResponse struct {
	Response []struct {
		League struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Country   string `json:"country"`
			Logo      string `json:"logo"`
			Flag      string `json:"flag"`
			Season    int    `json:"season"`
			Standings [][]struct {
				Rank int `json:"rank"`
				Team struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"team"`
				Points      int    `json:"points"`
				GoalsDiff   int    `json:"goalsDiff"`
				Group       string `json:"group"`
				Form        string `json:"form"`
				Status      string `json:"status"`
				Description string `json:"description"`
				All         struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"all"`
				Home struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"home"`
				Away struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"away"`
				Update string `json:"update"`
			} `json:"standings"`
		} `json:"league"`
	} `json:"response"`
}

func GetStanding(leagueId, season int) StandingsResponse {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/standings?season=%d&league=%d", season, leagueId)

	response := rapidRequest(url)

	var result StandingsResponse
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Fatalf("Error unmarshalling json: %v\n", err)
	}

	return result
}
