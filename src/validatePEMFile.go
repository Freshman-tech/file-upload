package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// ValidatePEM will check for pem file validity
func ValidatePEM(f string) {
	certFile := f
	certPEM, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatal("failed read certificatee file" + err.Error())
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(certPEM))
	if !ok {
		log.Fatal("failed to parse certificate, file will be removed")
		err := os.Remove(certFile)
		if err != nil {
			log.Error("could not delete file" + err.Error())
			return
		}
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		log.Fatal("Faile to parse certificate PEM" + err.Error())
		return
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal("failed to parse certificate" + err.Error())
		//fmt.Errorf("failed to parse certificate: %v", err.Error())
		return

	}

	opts := x509.VerifyOptions{
		// DNSName: name,
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		log.Fatal("failed to verify certificate" + err.Error())
		// fmt.Errorf("failed to verify certificate: %v", err.Error())
		return
	}

	// print cert infos
	certIssuer := cert.Issuer
	certCN := cert.Subject.CommonName
	certEndDate := cert.NotAfter.String()

	log.Info("successfully validated certificate for issuer: [", certIssuer, "] Common Name: [", certCN, "] valid until [", certEndDate, "]")
	return

}
