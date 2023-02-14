package resources

import (
	"encoding/json"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

func updateTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var track repository.Track
	if err := json.NewDecoder(r.Body).Decode(&track); err == nil {
		if id == track.Id {
			if n := repository.Update(track); n > 0 {
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(track); n > 0 {
				w.WriteHeader(201) /* Created */
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		w.WriteHeader(400) /* Bad Request */
	}
}

func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if track, num_found := repository.Read(id); num_found > 0 {
		d := repository.Track{Id: track.Id, Audio: track.Audio}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if num_found == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func listAllTrackIds(w http.ResponseWriter, r *http.Request) {
	if tracks, num_found := repository.ListAllIds(); num_found > 0 {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(tracks)
	} else if num_found == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if n := repository.Delete(id); n > 0 {
		w.WriteHeader(204) /* No Content */
	} else if n == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	/* Store */
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Document */
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	// List all tracks.
	r.HandleFunc("/tracks", listAllTrackIds).Methods("GET")
	// Delete a track.
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")
	return r
}
