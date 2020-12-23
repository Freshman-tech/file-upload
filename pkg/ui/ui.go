package ui

import (
	"net/http"

	"github.com/dadez/file-upload/pkg/common"
	log "github.com/sirupsen/logrus"
)

var (
	staticFilesPath string
)

// WebHandler serves the index.html file
func WebHandler(w http.ResponseWriter, r *http.Request) {

	// check for static files path
	staticFilesPath = common.GetEnv("STATIC_FILES_PATH", "./web") + "/"
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, staticFilesPath+"index.html")
	log.Debug("main page access")
}
