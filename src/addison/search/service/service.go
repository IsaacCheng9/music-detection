package service

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AudDMusic/audd-go"
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

func SearchAudDTracksAPI(base64Audio string) string {
	client := audd.NewClient(getAPIToken())
	url := "https://audd.tech/example.mp3"
	additionalParams := map[string]string{"audio": base64Audio}
	result, err := client.Recognize(url, "apple_music,spotify", additionalParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s - %s.\nTimecode: %s, album: %s. ℗ %s, %s\n\n"+
		"Listen: %s\nOr directly on:\n- Apple Music: %s, \n- Spotify: %s",
		result.Artist, result.Title, result.Timecode, result.Album,
		result.Label, result.ReleaseDate, result.SongLink,
		result.AppleMusic.URL, result.Spotify.ExternalUrls.Spotify)
	return result.Title
}
