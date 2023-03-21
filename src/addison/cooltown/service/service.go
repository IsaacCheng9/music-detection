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
				decodedBody := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&decodedBody); err == nil {
					if audio, err := getAudioFromBody(decodedBody); err == nil {
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

func getAudioFromBody(body map[string]interface{}) (string, error) {
	if audio, ok := body["Audio"].(string); ok {
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
				decodedBody := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&decodedBody); err == nil {
					if id, err := getIdFromBody(decodedBody); err == nil {
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

func getIdFromBody(body map[string]interface{}) (string, error) {
	if id, ok := body["Id"].(string); ok {
		return replaceSpacesWithPlusSymbols(id), nil
	}
	return "", errors.New("getIdFromBody")
}

// Replace spaces in the title with plus signs, as the AudD API returns a title
// with spaces.
func replaceSpacesWithPlusSymbols(id string) string {
	return strings.ReplaceAll(id, " ", "+")
}
