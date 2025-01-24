package threads

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"rest_imageboard/internal/storage"
	"rest_imageboard/internal/storage/query"

	"github.com/go-chi/chi/v5"
)

func Get(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thrs, err := query.GetThreads(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(thrs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func Save(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thr := query.Thread{}
		err := json.NewDecoder(r.Body).Decode(&thr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = query.SaveThread(db, thr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func Delete(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thr_name := chi.URLParam(r, "name")
		err := query.DeleteThread(db, query.Thread{Name: thr_name})
		if err == storage.ErrThreadNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
