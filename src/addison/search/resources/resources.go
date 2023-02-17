package resources

import (
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()
	/* controller */
	r.HandleFunc("/search", service.SearchTrack).Methods("POST")
	return r
}
