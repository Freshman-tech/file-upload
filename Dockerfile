## go application
FROM golang:1.14-alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /build
## We copy everything in the root directory
## into our /build directory
ADD src/* /build/
## We specify that we now wish to execute 
## any further commands inside our /build
## directory
WORKDIR /build
## we run go build to compile the binary
## executable of our Go program
RUN CGO_ENABLED=0 GOOS=linux go build -a -o fileserver .

# New stage for run app
#FROM alpine:3.11
FROM scratch

# get builded app from previous stage
COPY --from=0 /build/fileserver .

# add static files
ADD static .

# executable
ENTRYPOINT [ "./fileserver" ]
### Our start command which kicks off
### our newly created binary executable
#CMD ["/app/main"]
