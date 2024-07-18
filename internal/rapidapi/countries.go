package rapidapi

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Response []country `json:"response"`
}

type country struct {
	Name string `json:"name"`
}

func GetCountries() []country {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/countries")

	apiResponse := rapidRequest(url)

	var result response
	err := json.Unmarshal(apiResponse, &result)
	if err != nil {
		log.Fatalf("Erorr unmarshalling json: %v\n", err)
	}

	return result.Response
}
