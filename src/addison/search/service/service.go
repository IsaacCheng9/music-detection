package service

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

// !IMPORTANT: Store the API token for the AudD API.
const TOKEN = "a17de6c97a4dcc228fc6e051a19fbea3"

func titleOf(t map[string]interface{}) (string, error) {
	if response, ok := t["result"].(map[string]interface{}); ok {
		if title, ok := response["title"].(string); ok {
			return title, nil
		}
	}
	return "", errors.New("titleOf")
}

func SearchAuddTracksAPI(base64Audio string) (string, error) {
	data := url.Values{
		"api_token": {TOKEN},
		"audio":     {base64Audio},
		"return":    {"apple_music,spotify"},
	}
	response, _ := http.PostForm("https://api.audd.io/", data)
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	t := map[string]interface{}{}
	if err := json.NewDecoder(response.Body).Decode(&t); err == nil {
		if title, err := titleOf(t); err == nil {
			return title, nil
		} else {
			// The track could not be recognised by the AudD API.
			return "", nil
		}
	} else {
		// The AudD API was unable to process the request.
		return "", err
	}
}
