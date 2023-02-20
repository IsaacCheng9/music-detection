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
			if title, err := service.SearchAudDTracksAPI(base64audio); err == nil && title != "" {
				u := map[string]interface{}{"Id": title}
				json.NewEncoder(w).Encode(u)
				// 200 OK - the track has been found.
				w.WriteHeader(200)
			} else if err != nil {
				// 500 Internal Server Error - the API was unable to process the
				// request.
				w.WriteHeader(500)
			} else {
				// 404 Not Found - the track could not be recognised.
				w.WriteHeader(404)
			}
		}
	} else {
		// 400 Bad Request - the request could not be decoded by the server
		// due to malformed syntax.
		w.WriteHeader(400)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
