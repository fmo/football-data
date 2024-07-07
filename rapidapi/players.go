package rapidapi

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
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
	Age         int32  `json:"age"`
	Nationality string `json:"nationality"`
	Photo       string `json:"photo"`
}

type Statistics struct {
	Team      Team      `json:"team"`
	GameStats GameStats `json:"games"`
}

type GameStats struct {
	Appearances int32  `json:"appearances"`
	Position    string `json:"position"`
}

type Team struct {
	Name string `json:"name"`
	ID   int32  `json:"id"`
}

type RapidApi struct {
	logger *logrus.Logger
}

func NewRapidApi(l *logrus.Logger) RapidApi {
	return RapidApi{
		logger: l,
	}
}

func (r RapidApi) GetPlayers(page int, season int, teamId int) PlayerResponse {
	url := fmt.Sprintf(
		"https://api-football-v1.p.rapidapi.com/v3/players?season=%d&team=%d&page=%d",
		season,
		teamId,
		page,
	)

	// remove page param from the request if its page 0
	if page == 0 {
		url = url[:len(url)-7]
	}

	r.logger.Debug("request url: ", url)

	response := rapidRequest(url)

	var result PlayerResponse
	if err := json.Unmarshal(response, &result); err != nil {
		log.Fatalf("Error unmarshalling json: %v\n", err)
	}

	// Extracting first three player names
	playerNames := make([]string, 0, 3)
	for i, playerDetail := range result.Response {
		if i >= 3 { // Only interested in the first three players
			break
		}
		playerNames = append(playerNames, playerDetail.Player.Name)
	}

	r.logger.WithFields(logrus.Fields{
		"currentPage":     result.Paging.Current,
		"totalPages":      result.Paging.Total,
		"playersCount":    len(result.Response),
		"firstThreeNames": playerNames,
	}).Debug("Rapid API response summary with player names")

	return result
}
