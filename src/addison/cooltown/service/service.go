package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetIdFromAudioFragment(audio string) (string, error) {
	client := &http.Client{}
	uri := "http://localhost:3001/search"
	jsonData := []byte(`{"Audio": audio}`)
	if req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData)); err == nil {
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == http.StatusOK {
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if title, err := titleOf(t); err == nil {
						return title, nil
					}
				}
			}
		}
	}
	return "", errors.New("Service")
}

func titleOf(t map[string]interface{}) (string, error) {
	if title, ok := t["Id"].(string); ok {
		return replaceSpacesInTitle(title), nil
	}
	return "", errors.New("titleOf")
}

// Replace spaces in the title with plus signs, as the AudD API returns a title
// with spaces.
func replaceSpacesInTitle(title string) string {
	return strings.ReplaceAll(title, " ", "+")
}
