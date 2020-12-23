package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/dadez/file-upload/pkg/auth"
	"github.com/dadez/file-upload/pkg/common"
	"github.com/dadez/file-upload/pkg/healthz"
	"github.com/dadez/file-upload/pkg/ui"
	"github.com/dadez/file-upload/pkg/upload"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	logLevel string
)

func init() {
	// configure logging

	//log.SetReportCaller(true)

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// configure log level, default to info
	logLevel = common.GetEnv("LOG_LEVEL", "info")
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
}

// use gorilla mux for serve http
func main() {
	log.Info("fileupload server ready")
	// run webserver
	r := mux.NewRouter()

	//public
	r.HandleFunc("/healthz", healthz.HealthzHandler).Methods("GET")

	//private
	r.HandleFunc("/", auth.BasicAuth(ui.WebHandler)).Methods("GET")
	r.HandleFunc("/upload", auth.BasicAuth(upload.UploadHandler)).Methods("POST")

	if err := http.ListenAndServe(":4500", r); err != nil {
		log.Fatal(err)
	}
}
