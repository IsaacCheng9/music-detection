package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if base64audio, ok := t["Audio"].(string); ok {
			if base64audio != "" {
				if title, err := service.SearchAuddTracksAPI(base64audio); err == nil && title != "" {
					u := map[string]interface{}{"Id": title}
					if err := json.NewEncoder(w).Encode(u); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
					}
					// 200 OK - the track has been found.
					w.WriteHeader(http.StatusOK)
				} else if err != nil {
					// 500 Internal Server Error - the API was unable to process
					// the request.
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					// 404 Not Found - the track could not be recognised.
					w.WriteHeader(http.StatusNotFound)
				}
			} else {
				// 404 Not Found - the track could not be recognised.
				w.WriteHeader(http.StatusNotFound)
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
