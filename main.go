package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/", HelloWorldHandler)

	appengine.Main()
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Warningf(r.Context(), "failed httpResponseWrite. err=%+v\n", err)
	}
}
