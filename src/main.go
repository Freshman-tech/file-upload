package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const maxUploadSize = 1024 * 1024 // 1MB

var (
	staticFilesPath      string
	uploadsDirectoryPath string
)

// configure logging
func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	// check for environment variables
	// check for static files path
	staticFilesPath = GetEnv("STATIC_FILES_PATH", ".") + "/"
	// check for uploads directory path
	uploadsDirectoryPath = GetEnv("UPLOADS_DIRECTORY_PATH", "uploads")
}

// GetEnv checks for existing environment variables
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// IndexHandler serves the index.html file
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	http.ServeFile(w, r, staticFilesPath+"index.html")
	log.Debug("main page access")
}

// use gorilla mux for serve http
func main() {
	log.Info("fileupload server ready")
	// run webserver
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/upload", BasicAuth(uploadHandler)).Methods("POST")

	if err := http.ListenAndServe(":4500", r); err != nil {
		log.Fatal(err)
	}
}
