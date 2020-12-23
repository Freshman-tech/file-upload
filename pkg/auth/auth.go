package auth

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dadez/file-upload/pkg/common"
	log "github.com/sirupsen/logrus"
)

var authFilesPath string

// Handler defines the incomming request
type Handler func(w http.ResponseWriter, r *http.Request)

// BasicAuth setting auth for desired routes
func BasicAuth(pass Handler) Handler {

	return func(w http.ResponseWriter, r *http.Request) {

		// ask for credentials if the Authorization header parse fails
		_, _, ok := r.BasicAuth()
		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "No basic auth present"}`))
			return
		}

		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			log.Error("Failed authorization access, no credentials provided")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		//if len(pair) != 2 || !validate(pair[0], pair[1]) {
		if len(pair) != 2 || !validateFromFile(pair[0], pair[1]) {
			log.Error("Failed authorization access, username and or password missmatch")
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

// only for testing purposes, should not be used in production
func validate(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}

// username and password should be defined in files
// on kubernetes this files will be mounted from a secret file
func validateFromFile(username, password string) bool {
	// check for static files path
	authFilesPath = common.GetEnv("AUTH_FILES_PATH", "./test") + "/"

	// get username from file
	user, err := os.Open(authFilesPath + "username")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = user.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	u, err := ioutil.ReadAll(user)
	log.Debug("username is: ", string(u))

	// get password from file
	pass, err := os.Open(authFilesPath + "password")
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
	log.Debug("password is: ", string(p))

	if username == string(u) && password == string(p) {
		log.Debug("login successful")
		return true
	}
	return false
}
