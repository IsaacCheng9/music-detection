package main

import (
	"log"
	"net/http"
	"tracks/repository"
	"tracks/resources"
)

func main() {
	repository.Init()
	// Drop the Tracks table so that we can start with a clean database.
	repository.Clear()
	repository.Create()
	log.Fatal(http.ListenAndServe(":3000", resources.Router()))
}
