package rapidapi

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func rapidRequest(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "api-football-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body
}
