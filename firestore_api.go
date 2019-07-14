package main

import (
	"net/http"

	"github.com/sinmetal/gcpmetadata"
	"google.golang.org/appengine/log"
)

func FirestoreHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectID, err := gcpmetadata.GetProjectID()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	fc, err := NewFirestoreClient(ctx, projectID)
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
