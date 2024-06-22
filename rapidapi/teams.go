package rapidapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type TeamsResponse struct {
	Response []TeamsSpecs `json:"response"`
}

type TeamsSpecs struct {
	Team team `json:"team"`
}

type team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func GetTeams(leagueId int) []TeamsSpecs {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?league=%d&season=%d", leagueId, 2023)

	response := rapidRequest(url)

	var result TeamsResponse
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Fatalf("Error unmarshalling json: %v\n", err)
	}

	return result.Response
}
