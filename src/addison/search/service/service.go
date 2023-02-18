package service

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Get the API token from token.txt.
func getAPIToken() string {
	file, err := os.Open("../token.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func titleOf(t map[string]interface{}) (string, error) {
	if response, ok := t["result"].(map[string]interface{}); ok {
		if title, ok := response["title"].(string); ok {
			return title, nil
		}
	}
	return "", errors.New("titleOf")
}

func SearchAudDTracksAPI(base64Audio string) (string, error) {
	//client := audd.NewClient(getAPIToken())
	//client := audd.NewClient("b30fbcc812b45d5379114bd1f430a32c")
	//url := "https://audd.tech/example.mp3"
	//additionalParams := map[string]string{"audio": base64Audio}
	//result, err := client.Recognize(url, "apple_music,spotify", additionalParams)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s - %s.\nTimecode: %s, album: %s. ℗ %s, %s\n\n"+
	//	"Listen: %s\nOr directly on:\n- Apple Music: %s, \n- Spotify: %s",
	//	result.Artist, result.Title, result.Timecode, result.Album,
	//	result.Label, result.ReleaseDate, result.SongLink,
	//	result.AppleMusic.URL, result.Spotify.ExternalUrls.Spotify)
	//return result.Title

	data := url.Values{
		"api_token": {"b30fbcc812b45d5379114bd1f430a32c"},
		"audio":     {base64Audio},
		"return":    {"apple_music,spotify"},
	}
	response, _ := http.PostForm("https://api.audd.io/", data)
	defer response.Body.Close()
	//body, _ := io.ReadAll(response.Body)
	//fmt.Println(string(body))
	t := map[string]interface{}{}
	if err := json.NewDecoder(response.Body).Decode(&t); err == nil {
		if title, err := titleOf(t); err == nil {
			return title, nil
		}
	}

	return "", errors.New("SearchAudDTracksAPI")
}
