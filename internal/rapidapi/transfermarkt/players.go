package transfermarkt

import (
	"encoding/json"
	"fmt"
	"github.com/fmo/football-data/internal/rapidapi"
	"log"
)

type Player struct {
	Height        string        `json:"height"`
	Foot          string        `json:"foot"`
	Injury        interface{}   `json:"injury"`
	Suspension    interface{}   `json:"suspension"`
	Joined        int64         `json:"joined"`
	ContractUntil int64         `json:"contractUntil"`
	Captain       bool          `json:"captain"`
	LastClub      interface{}   `json:"lastClub"`
	IsLoan        interface{}   `json:"isLoan"`
	WasLoan       interface{}   `json:"wasLoan"`
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Image         string        `json:"image"`
	ImageLarge    interface{}   `json:"imageLarge"`
	ImageSource   string        `json:"imageSource"`
	ShirtNumber   interface{}   `json:"shirtNumber"`
	Age           int           `json:"age"`
	DateOfBirth   int64         `json:"dateOfBirth"`
	HeroImage     string        `json:"heroImage"`
	IsGoalkeeper  bool          `json:"isGoalkeeper"`
	Positions     Positions     `json:"positions"`
	Nationalities []Nationality `json:"nationalities"`
	MarketValue   MarketValue   `json:"marketValue"`
}

type Positions struct {
	First  Position  `json:"first"`
	Second *Position `json:"second"`
	Third  *Position `json:"third"`
}

type Position struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Group     string `json:"group"`
}

type Nationality struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type MarketValue struct {
	Value       int         `json:"value"`
	Currency    string      `json:"currency"`
	Progression interface{} `json:"progression"`
}

type Data struct {
	Players []Player `json:"data"`
}

func GetPlayers() []byte {
	url := fmt.Sprintf("https://transfermarkt-db.p.rapidapi.com/v1/clubs/squad?season_id=%d&locale=UK&club_id=%d",
		2024,
		12321,
	)

	response := rapidapi.RapidRequest(url)

	var playerResponse Data
	if err := json.Unmarshal(response, &playerResponse); err != nil {
		log.Fatalf("error unmarshalling json: %v\n", err)
	}

	for _, p := range playerResponse.Players {
		fmt.Println(p.Name)
	}
	return nil
}
