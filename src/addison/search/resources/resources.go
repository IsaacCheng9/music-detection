package resources

import (
	"encoding/json"
	"net/http"
	"search/service"

	"github.com/gorilla/mux"
)

func searchTrack(w http.ResponseWriter, r *http.Request) {
	//base64Audio := r.FormValue("Audio")
	// Should work with the script:
	// ID="~Everybody+(Backstreet ’s+Back)+(Radio+Edit)" AUDIO=‘base64 -i "$ID".wav‘
	// RESOURCE=localhost:3001/search
	// echo "{ \"Audio\":\"$AUDIO\" }" > input curl -v -X POST -d @input $RESOURCE
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

	//vars := mux.Vars(r)
	//base64Audio := vars["Audio"]
	//fmt.Println(base64Audio)
	//result := service.SearchAudDTracksAPI(base64Audio)
	//w.WriteHeader(200) /* OK */
	//json.NewEncoder(w).Encode(result)
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/search", searchTrack).Methods("POST")
	return r
}
