# How to process file uploads in Go

This repo contains the complete code used in [this Freshman
tutorial](https://freshman.tech/file-upload-golang/). Clone this repo to your
computer and run `go run main.go` to start the server on PORT 4500.


# Goal

The purpose of this program is to be able to upload certificate files in a PEM format for later use as sidecar container together with https://github.com/joe-elliott/cert-exporter


# Usage

## start locally

```bash
cd file-upload # move directory to this project
cd src # move to directory containg go files
export STATIC_FILES_PATH="../static" # tell where the index.html is stored, fallback to '.' if not defined
export AUTH_FILES_PATH="../static" # tell where the credentials files are stored, fallback to '.' if not defined
export UPLOADS_DIRECTORY_PATH="../uploads" # tell where the uploaded files should be stored, fallback to 'uploads' if not defined

# run application
go run main.go auth.go upload.go validatePEMFile.go
```

you should see the following output: `fileupload server ready`


## Upload a cerfificate with curl

| using test credentials  
| On production (kubernetes), this files should be overriden

```bash
cd file-upload # move directory to this project
curl -X POST -u user:secret -F file=@test/github.crt.pem http://localhost:4500/upload
```

## Upload mulitples files with curl

just repeat the `-F file=@/path/to/file.pem` part in your command