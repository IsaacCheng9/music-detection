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
			if title, err := service.SearchAudDTracksAPI(base64audio); err == nil {
				u := map[string]interface{}{"Id": title}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(u)
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
