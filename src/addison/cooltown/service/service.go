package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAudioFromId(id string) (string, error) {
	client := &http.Client{}
	uri := "http://localhost:3000/tracks/" + id
	fmt.Println("2a")
	fmt.Println(uri)
	if req, err := http.NewRequest("GET", uri, nil); err == nil {
		fmt.Println("2b")
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == http.StatusOK {
				fmt.Println("2c")
				t := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&t); err == nil {
					if audio, err := audioOf(t); err == nil {
						return audio, nil
					}
				}
			}
		}
	}
	return "", errors.New("Service")
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
	fmt.Println("1a")
	//jsonData := []byte("{'Audio':" + audio + "}")
	jsonData, _ := json.Marshal(map[string]string{"Audio": audio})
	if req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData)); err == nil {
		fmt.Println("1b")
		if rsp, err := client.Do(req); err == nil {
			if rsp.StatusCode == http.StatusOK {
				fmt.Println("1c")
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
