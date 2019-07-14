package main

import (
	"net/http"

	"google.golang.org/appengine/log"
)

func FirestoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fc, err := NewFirestoreClient(ctx, ProjectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	s, err := NewFirestoreStore(ctx, fc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	d := &Sample{}
	d, err = s.Create(ctx, d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(d.ID))
	if err != nil {
		log.Warningf(r.Context(), "failed httpResponseWrite. err=%+v\n", err)
	}
}
