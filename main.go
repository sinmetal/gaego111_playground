package main

import (
	"fmt"
	"net/http"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/sinmetal/gcpmetadata"
	"go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var ProjectID string

func main() {
	projectID, err := gcpmetadata.GetProjectID()
	if err != nil {
		panic(err)
	}
	ProjectID = projectID
	fmt.Printf("ProjectID=%s\n", projectID)

	if gcpmetadata.OnGCP() {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID: projectID,
		})
		if err != nil {
			panic(err)
		}
		trace.RegisterExporter(exporter)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/firestore/", FirestoreHandler)
	mux.HandleFunc("/", HelloWorldHandler)

	http.Handle("/", &ochttp.Handler{
		Handler:     mux,
		Propagation: &propagation.HTTPFormat{},
	})

	appengine.Main()
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Warningf(r.Context(), "failed httpResponseWrite. err=%+v\n", err)
	}
}
