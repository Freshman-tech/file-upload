package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type handler func(w http.ResponseWriter, r *http.Request)

// BasicAuth setting auth for desired routes
func BasicAuth(pass handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		//if len(pair) != 2 || !validate(pair[0], pair[1]) {
		if len(pair) != 2 || !validateCredentials(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

func validate(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}

func validateCredentials(username, password string) bool {
	// get username from file
	user, err := os.Open("username")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = user.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	u, err := ioutil.ReadAll(user)

	// get password from file
	pass, err := os.Open("password")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = pass.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// check if username and password match
	p, err := ioutil.ReadAll(pass)

	if username == string(u) && password == string(p) {
		return true
		log.Debug("login successfull")
	}
	return false
}
