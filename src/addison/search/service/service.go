package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// Store the API token for the AudD API.
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
	defer response.Body.Close()
	t := map[string]interface{}{}
	if err := json.NewDecoder(response.Body).Decode(&t); err == nil {
		if title, err := titleOf(t); err == nil {
			return title, nil
		}
	}

	return "", errors.New("SearchAudDTracksAPI")
}
