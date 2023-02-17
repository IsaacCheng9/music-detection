package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	base64Audio := r.FormValue("Audio")
	result := service.SearchAudDTracksAPI(base64Audio)
	w.WriteHeader(200) /* OK */
	json.NewEncoder(w).Encode(result)
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* controller */
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
