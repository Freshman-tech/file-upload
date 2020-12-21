package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Healthz Handler for use in kubernetes
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	started := time.Now()
	duration := time.Since(started)
	if duration.Seconds() > 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		log.Error("health check takes too long: ", duration.String())
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
		log.Debug("healthz check tooks: ", duration.String())
	}
}
