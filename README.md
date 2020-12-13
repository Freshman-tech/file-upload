# How to process file uploads in Go

This repo contains the complete code used in [this Freshman
tutorial](https://freshman.tech/file-upload-golang/). Clone this repo to your
computer and run `go run main.go` to start the server on PORT 4500.


# Usage

## curl

uploading images works well with curl
```bash
curl -X POST -F "type=file" -F "file=@/path/to/file.png" http://localhost:4500/upload
```

uploading pdf files not, even setting the type in path to file
```bash
curl -X POST -F "type=file" -F "file=@/path/to/file.pdf;type=application/pdf" http://localhost:4500/upload
```
the file is seen as plain text `level=debug msg="File type is: text/plain; charset=utf-8"
`
what does the browser doing different as curl ?


# Auth

## dummy auth implemented for /auth, see auth.go

```bash
curl -u test:test localhost:4500/auth
```