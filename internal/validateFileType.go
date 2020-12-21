package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	// libs for managing mime types, does not recognize x509
	"github.com/h2non/filetype"
)

var pemType = filetype.NewType("pem", "application/x-x509-ca-cert")

func pemMatcher(buf []byte) bool {
	return len(buf) > 1 && buf[0] == 0x01 && buf[1] == 0x02
}

// ValidateFileType check mime Type based on content of filetype package
func ValidateFileType(f []byte) {
	// Check if file is supported by extension
	// Register the new matcher and its type
	filetype.AddMatcher(pemType, pemMatcher)

	// Check if the new type is supported by extension
	if filetype.IsSupported("pem") {
		fmt.Println("New supported type: pem")
	}

	// Check if the new type is supported by MIME
	if filetype.IsMIMESupported("application/x-x509-ca-cert") {
		fmt.Println("New supported MIME type: application/x-x509-ca-cert")
	}

	kind, _ := filetype.Match(f)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type extension", kind.Extension)
	} else {
		fmt.Printf("File type matched: %s\n", kind.Extension)
	}

}

// ValidateContentType check mime Type based on content of http package
func ValidateContentType(f []byte) {

	// ensure uploaded file is a certificate in the pem format
	// todo how to recognize pplication/x-x509-ca-cert mime Type
	contentType := http.DetectContentType(f)
	log.Debug("File type is: ", contentType)
	// if contentType != "application/x-x509-ca-cert" && contentType != "application/pdf" {
	if contentType != "application/x-x509-ca-cert" {
		// http.Error(w, "The provided file format is not allowed.", http.StatusBadRequest)
		// log.Error("Uploading file [", fileHeader.Filename, "] of type [", contentType, "] failed, file type not allowed")
		log.Error("Uploading file [] of type [", contentType, "] failed, file type not allowed")
		return
	}
}

// ValidateExtension checks for file extension
// func ValidateExtension(f string) {
// 	mime.AddExtensionType(".pem", "application/x-x509-ca-cert")
// 	mime.AddExtensionType(".crt", "application/x-x509-ca-cert")

// 	if filepath.Ext(f) != ".pem" {
// 		log.Error("only .pem are allowed")
// 		return
// 	}

// }
