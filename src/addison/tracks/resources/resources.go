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
				// 204 No Content - the track already existed and has been
				// updated successfully.
				w.WriteHeader(204)
			} else if n := repository.Insert(track); n > 0 {
				// 201 Created - the track has been created successfully.
				w.WriteHeader(201)
			} else {
				w.WriteHeader(500)
			}
		} else {
			// 400 Bad Request - the id in the URL does not match the id in the
			// request body.
			w.WriteHeader(400)
		}
	} else {
		// 400 Bad Request - the request body could not be decoded.
		w.WriteHeader(400)
	}
}

func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if track, numFound := repository.Read(id); numFound > 0 {
		d := repository.Track{Id: track.Id, Audio: track.Audio}
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(d)
	} else if numFound == 0 {
		w.WriteHeader(404) /* Not Found */
	} else {
		w.WriteHeader(500) /* Internal Server Error */
	}
}

func listAllTrackIds(w http.ResponseWriter, _ *http.Request) {
	if tracks, numFound := repository.ListAllIds(); numFound > 0 {
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(tracks)
	} else if numFound == 0 {
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
