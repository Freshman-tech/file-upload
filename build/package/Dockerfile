FROM golang:1.14-alpine

RUN mkdir /build
ADD cmd /build/cmd
ADD pkg /build/pkg
ADD go.mod /build/
WORKDIR /build
RUN ls -l
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o fileupload ./cmd/app

FROM alpine:3.12
#FROM scratch

RUN addgroup -g 1000 app && \
    adduser -u 1000 -h /app -G app -S app

WORKDIR /app
USER app

# get builded app from previous stage
COPY --from=0 /build/fileupload .

# add static files
ADD web web
ADD test test

# executable
CMD ["/app/fileupload"]
