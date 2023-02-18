package resources

import (
	"cooltown/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getTrackFromFragment(w http.ResponseWriter, r *http.Request) {
	t := map[string]interface{}{}
	var id string
	fmt.Println("1")
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {
		if base64Audio, ok := t["Audio"].(string); ok {
			if title, err := service.GetIdFromAudioFragment(base64Audio); err == nil {
				id = title
			} else {
				fmt.Println("1 - Track 404")
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println("2")
	if audio, err := service.GetAudioFromId(id); err == nil {
		u := map[string]interface{}{"Audio": audio}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
		return
	} else {
		fmt.Println("3 - Audio 404")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/cooltown", getTrackFromFragment).Methods("POST")
	return r
}
