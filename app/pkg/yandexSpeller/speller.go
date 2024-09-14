package yandexspeller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type yandexSpellerAPIReponse struct {
	Code     int    `json:"code"`
	Position int    `json:"pos"`
	Row      int    `json:"row"`
	Column   int    `json:"col"`
	Length   int    `json:"len"`
	Word     string   `json:"word"`
	S        []string `json:"s"`
}

var client = &http.Client{}

// TODO: Add all parameters
func MakeRequst(text string) ([]yandexSpellerAPIReponse, error) {
	logrus.Tracef("making request to yandex speller API")

	req := &http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Scheme:   "https",
			Host:     "speller.yandex.net",
			Path:     "/services/spellservice.json/checkText",
			RawQuery: url.Values{"text": {text}}.Encode(),
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		return nil, fmt.Errorf("failed to make request, status code: %d, response body: %s", resp.StatusCode, body)
	}

	var spellerResponse []yandexSpellerAPIReponse
	json.NewDecoder(resp.Body).Decode(&spellerResponse)

	return spellerResponse, nil
}
