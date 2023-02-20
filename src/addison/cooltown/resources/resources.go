package resources

import (
	"cooltown/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getTrackFromFragment(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{}
	var id string
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if base64Audio, ok := t["Audio"].(string); ok {
			if title, err := service.GetIdFromAudioFragment(base64Audio); err == nil && title != "" {
				id = title
			} else if err != nil {
				// 500 Internal Server Error - the server was unable to process
				// the request.
				w.WriteHeader(500)
				return
			} else {
				// 404 Not Found - no matching track and ID for the audio
				// fragment was found.
				w.WriteHeader(404)
				return
			}
		}
	} else {
		// 400 Bad Request - the request could not be decoded by the server due
		// to malformed syntax.
		w.WriteHeader(400)
	}

	if audio, err := service.GetAudioFromId(id); err == nil && audio != "" {
		u := map[string]interface{}{"Audio": audio}
		json.NewEncoder(w).Encode(u)
		// 200 OK - the matching track for the audio fragment has been found.
		w.WriteHeader(200)
	} else if err != nil {
		// 500 Internal Server Error - the server was unable to process the
		// request.
		w.WriteHeader(500)
	} else {
		// 404 Not Found - no matching track for the audio fragment was found.
		w.WriteHeader(404)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", getTrackFromFragment).Methods("POST")
	return r
}
