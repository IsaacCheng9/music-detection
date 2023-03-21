package resources

import (
	"cooltown/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getTrackFromFragment(w http.ResponseWriter, r *http.Request) {
	decodedTrack := map[string]interface{}{}
	var id string

	// Part 1: Get the ID of the track from the audio fragment using the Search
	// microservice.
	if err := json.NewDecoder(r.Body).Decode(&decodedTrack); err == nil {
		if base64Audio, ok := decodedTrack["Audio"].(string); ok {
			if base64Audio != "" {
				if title, err := service.GetIdFromAudioFragment(base64Audio); err == nil && title != "" {
					id = title
				} else if err != nil {
					// 500 Internal Server Error - the server was unable to
					// process the request.
					w.WriteHeader(http.StatusInternalServerError)
					return
				} else {
					// 404 Not Found - no matching track and ID for the audio
					// fragment was found.
					w.WriteHeader(http.StatusNotFound)
					return
				}
			} else {
				// 404 Not Found - no matching track and ID for the audio
				// fragment was found.
				w.WriteHeader(http.StatusNotFound)
				return
			}
		} else {
			// 400 Bad Request - the request could not be decoded by the server
			// as the 'Audio' field is missing.
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		// 400 Bad Request - the request could not be decoded by the server due
		// to malformed syntax.
		w.WriteHeader(http.StatusBadRequest)
	}

	// Part 2: Get the audio fragment from the ID using the Tracks microservice.
	if audio, err := service.GetAudioFromId(id); err == nil && audio != "" {
		encodedAudio := map[string]interface{}{"Audio": audio}
		if err := json.NewEncoder(w).Encode(encodedAudio); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// 200 OK - the matching track for the audio fragment has been found.
		w.WriteHeader(http.StatusOK)
	} else if err != nil {
		// 500 Internal Server Error - the server was unable to process the
		// request as the database is down.
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// 404 Not Found - no matching track for the audio fragment was found in
		// the database.
		w.WriteHeader(http.StatusNotFound)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", getTrackFromFragment).Methods("POST")
	return r
}
