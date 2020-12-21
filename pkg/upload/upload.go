package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dadez/file-upload/pkg/common"
	"github.com/dadez/file-upload/pkg/validate"
	log "github.com/sirupsen/logrus"
)

const (
	maxUploadSize = 1024 * 1024 // 1MB
)

// UploadHandler manages the upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("method is: ", r.Method)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Error("Method ", r.Method, " not allowed")
		return
	}

	// 32 MB is the default used by FormFile
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get a reference to the fileHeaders
	files := r.MultipartForm.File["file"]
	log.Debug("files are", files)

	for _, fileHeader := range files {
		log.Debug("check for upload file size ", fileHeader.Filename)
		if fileHeader.Size > maxUploadSize {
			http.Error(w, fmt.Sprintf("The uploaded file is too big: %s. Please use a file less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			log.Error("Upload failed, ", fileHeader.Filename, " is too big (", fileHeader.Size, ")")
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// check for file extension we only allow pem files
		if filepath.Ext(fileHeader.Filename) != ".pem" {
			http.Error(w, "The provided file format is not allowed.", http.StatusBadRequest)
			log.Error("uploading file ", fileHeader.Filename, " failed, only .pem are allowed")
			return
		}

		// todo: check what does seek do
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Debug("creating upload directory")
		uploadsDirectoryPath := common.GetEnv("UPLOADS_DIRECTORY_PATH", "uploads")

		err = os.MkdirAll(uploadsDirectoryPath, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		// create empty file on disk
		f, err := os.Create(uploadsDirectoryPath + "/" + fileHeader.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		// write file
		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Error("Error writing file")
			return
		}

		// validate the certificate
		validate.ValidatePEM(uploadsDirectoryPath + "/" + fileHeader.Filename)
		// // is there a way to get this informations back from validatePEMF function ?
		// certIssuer := Cert.Issuer
		// certCN := Cert.Subject.CommonName
		// certEndDate := Cert.NotAfter.String()

		// upload successful
		fmt.Fprintf(w, "successfully uploaded %s\n", fileHeader.Filename)
		// log.Info("Successful upload of file [", fileHeader.Filename, "] from issuer [", CertIssuer, "]")
		log.Info("Successfully uploaded file [", fileHeader.Filename, "]")
		//log.Debug("  certificate Common Name: [", CertCN, "] certificate valid until: [", CertEndDate, "]")
	}
}
