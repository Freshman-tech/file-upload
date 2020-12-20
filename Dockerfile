FROM golang:1.14-alpine

RUN mkdir /build
ADD src/* /build/
WORKDIR /build
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o fileserver .

FROM alpine:3.12
#FROM scratch

RUN addgroup -g 1000 app && \
    adduser -u 1000 -h /app -G app -S app

WORKDIR /app
USER app

# get builded app from previous stage
COPY --from=0 /build/fileserver .

# add static files
ADD static .

# executable
CMD ["/app/fileserver"]
