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
				w.WriteHeader(http.StatusNoContent)
			} else if n := repository.Insert(track); n > 0 {
				// 201 Created - the track has been created successfully.
				w.WriteHeader(http.StatusCreated)
			} else {
				// 500 Internal Server Error - the database is not available.
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			// 400 Bad Request - the id in the URL does not match the id in the
			// request body.
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		// 400 Bad Request - the request body could not be decoded.
		w.WriteHeader(http.StatusBadRequest)
	}
}

func readTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if track, numFound := repository.Read(id); numFound > 0 {
		d := repository.Track{Id: track.Id, Audio: track.Audio}
		if err := json.NewEncoder(w).Encode(d); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// 200 OK - the track has been returned successfully.
		w.WriteHeader(http.StatusOK)
	} else if numFound == 0 {
		// 404 Not Found - the track does not exist.
		w.WriteHeader(http.StatusNotFound)
	} else {
		// 500 Internal Server Error - the database is not available.
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func listAllTrackIds(w http.ResponseWriter, _ *http.Request) {
	if tracks, numFound := repository.ListAllIds(); numFound >= 0 {
		if err := json.NewEncoder(w).Encode(tracks); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// 200 OK - the list of tracks has been returned successfully.
		w.WriteHeader(http.StatusOK)
	} else {
		// 500 Internal Server Error - the database is not available.
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func deleteTrack(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if n := repository.Delete(id); n > 0 {
		// 204 No Content - the track has been deleted successfully.
		w.WriteHeader(http.StatusNoContent)
	} else if n == 0 {
		// 404 Not Found - the track does not exist.
		w.WriteHeader(http.StatusNotFound)
	} else {
		// 500 Internal Server Error - the database is not available.
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	r.HandleFunc("/tracks", listAllTrackIds).Methods("GET")
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")
	return r
}
