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
					if audio, err := getAudioFromBody(t); err == nil {
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

func getAudioFromBody(t map[string]interface{}) (string, error) {
	if audio, ok := t["Audio"].(string); ok {
		return audio, nil
	}
	return "", errors.New("getAudioFromBody")
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
					if id, err := getIdFromBody(t); err == nil {
						return id, nil
					}
				}
			} else if rsp.StatusCode == 404 {
				return "", nil
			}
		}
	}
	return "", errors.New("GetIdFromAudioFragment")
}

func getIdFromBody(t map[string]interface{}) (string, error) {
	if id, ok := t["Id"].(string); ok {
		return replaceSpacesWithPlusSymbols(id), nil
	}
	return "", errors.New("getIdFromBody")
}

// Replace spaces in the title with plus signs, as the AudD API returns a title
// with spaces.
func replaceSpacesWithPlusSymbols(id string) string {
	return strings.ReplaceAll(id, " ", "+")
}
