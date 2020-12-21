package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// web ui Handler serves the index.html file
func uiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, staticFilesPath+"index.html")
	log.Debug("main page access")
}
