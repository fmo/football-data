package rapidapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type PlayerResponse struct {
	Paging   Paging          `json:"paging"`
	Response []PlayerDetails `json:"response"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type PlayerDetails struct {
	Player     player       `json:"player"`
	Statistics []Statistics `json:"statistics"`
}

type player struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Photo       string `json:"photo"`
}

type Statistics struct {
	Team      Team      `json:"team"`
	GameStats GameStats `json:"games"`
}

type GameStats struct {
	Appearances int    `json:"appearances"`
	Position    string `json:"position"`
}

type Team struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func GetPlayers(page int, season int, teamId int) PlayerResponse {
	url := fmt.Sprintf(
		"https://api-football-v1.p.rapidapi.com/v3/players?season=%d&team=%d&page=%d",
		season,
		teamId,
		page,
	)

	response := rapidRequest(url)

	var result PlayerResponse
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatalf("Error unmarshalling json: %v\n", err)
	}

	return result
}
