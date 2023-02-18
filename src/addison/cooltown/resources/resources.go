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
			if title, err := service.GetIdFromAudioFragment(base64Audio); err == nil {
				id = title
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	if audio, err := service.GetAudioFromId(id); err == nil {
		u := map[string]interface{}{"Audio": audio}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
		return
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", getTrackFromFragment).Methods("POST")
	return r
}
