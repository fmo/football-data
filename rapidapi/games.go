package rapidapi

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type FixtureDetails struct {
	Teams   teams   `json:"teams"`
	Fixture Fixture `json:"fixture"`
	Score   Score   `json:"score"`
	League  League  `json:"league"`
}

type League struct {
	Round string `json:"round"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

type Fixture struct {
	Id        int    `json:"id"`
	Status    Status `json:"status"`
	Timestamp int    `json:"timestamp"`
	GameDate  string `json:"date"`
}

type Status struct {
	Short string `json:"short"`
}

type Score struct {
	FullTime FullTime `json:"fullTime"`
}

type FullTime struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type teams struct {
	Home TeamDetails `json:"home"`
	Away TeamDetails `json:"away"`
}

type TeamDetails struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner *bool  `json:"winner,omitempty"` // Use pointer to bool to allow for null
}

type FixtureResponse struct {
	Response []FixtureDetails `json:"response"`
}

func GetGames(rp map[string]int) []FixtureDetails {
	url := fmt.Sprintf(
		"https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%d&season=%d",
		rp["leagueId"],
		rp["season"],
	)

	if teamId, ok := rp["teamId"]; ok {
		url = fmt.Sprintf("%s&team=%d", url, teamId)
	}

	response := rapidRequest(url)

	var result FixtureResponse
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatalf("Error unmarshalling json: %v\n", err)
	}

	return result.Response
}
