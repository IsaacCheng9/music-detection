package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	decodedTrack := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&decodedTrack); err == nil {
		if base64audio, ok := decodedTrack["Audio"].(string); ok {
			if base64audio != "" {
				if title, err := service.SearchAuddRecognitionAPI(base64audio); err == nil && title != "" {
					trackId := map[string]interface{}{"Id": title}
					// 200 OK - the track has been found.
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(trackId)
				} else if err != nil {
					// 500 Internal Server Error - the API was unable to process
					// the request.
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					// 404 Not Found - the track could not be recognised.
					w.WriteHeader(http.StatusNotFound)
				}
			} else {
				// 400 Bad Request - no audio was found in the file provided.
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			// 400 Bad Request - the request could not be decoded by the server
			// as the 'Audio' field is missing.
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		// 400 Bad Request - the request could not be decoded by the server
		// due to malformed syntax.
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
