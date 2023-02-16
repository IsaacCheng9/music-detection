package main

import (
	"bufio"
	"fmt"
	"io"
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
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	data := url.Values{
		"url":    {"https://audd.tech/example.mp3"},
		"return": {"apple_music,spotify"},
		"api_token": {getAPIToken()},
	}
	response, _ := http.PostForm("https://api.audd.io/", data)
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
