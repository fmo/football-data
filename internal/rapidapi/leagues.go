package rapidapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type leaguesResponse struct {
	Response []leagueSpecs `json:"response"`
}

type leagueSpecs struct {
	League league `json:"league"`
}

type league struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetLeagues() []leagueSpecs {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues")

	response := rapidRequest(url)

	var result leaguesResponse
	err := json.Unmarshal(response, &result)
	if err != nil {
		log.Fatalf("Erorr unmarshalling json: %v\n", err)
	}

	return result.Response
}
