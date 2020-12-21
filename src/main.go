package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const (
	maxUploadSize = 1024 * 1024 // 1MB
)

var (
	staticFilesPath      string
	authFilesPath        string
	uploadsDirectoryPath string
	logLevel             string
)

func init() {
	// configure logging

	//log.SetReportCaller(true)

	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// configure log level, default to info
	logLevel = GetEnv("LOG_LEVEL", "info")
	lvl := strings.ToLower(logLevel)

	switch lvl {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}

	// check for environment variables
	// check for static files path
	staticFilesPath = GetEnv("STATIC_FILES_PATH", ".") + "/"
	// check for static files path
	authFilesPath = GetEnv("AUTH_FILES_PATH", ".") + "/"
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

// use gorilla mux for serve http
func main() {
	log.Info("fileupload server ready")
	log.Debug("auth files under ", authFilesPath)
	// run webserver
	r := mux.NewRouter()

	r.HandleFunc("/", uiHandler).Methods("GET")
	r.HandleFunc("/healthz", healthzHandler).Methods("GET")
	r.HandleFunc("/upload", BasicAuth(uploadHandler)).Methods("POST")

	if err := http.ListenAndServe(":4500", r); err != nil {
		log.Fatal(err)
	}
}
