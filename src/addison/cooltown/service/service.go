package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetAudioFromId(id string) (string, error) {
	client := &http.Client{}
	uri := "http://localhost:3000/tracks/" + id
	if req, err := http.NewRequest("GET", uri, nil); err == nil {
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == 200 {
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if audio, err := audioOf(t); err == nil {
						return audio, nil
					}
				}
			} else if rsp.StatusCode == 404 {
				return "", nil
			}
		}
	}
	return "", errors.New("GetAudioFromId")
}

func audioOf(t map[string]interface{}) (string, error) {
	if audio, ok := t["Audio"].(string); ok {
		return audio, nil
	}
	return "", errors.New("audioOf")
}

func GetIdFromAudioFragment(audio string) (string, error) {
	client := &http.Client{}
	uri := "http://localhost:3001/search"
	jsonData, _ := json.Marshal(map[string]string{"Audio": audio})
	if req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData)); err == nil {
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == 200 {
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if title, err := titleOf(t); err == nil {
						return title, nil
					}
				}
			} else if rsp.StatusCode == 404 {
				return "", nil
			}
		}
	}
	return "", errors.New("GetIdFromAudioFragment")
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
