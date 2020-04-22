############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
ENV CGO_ENABLED=0
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
#RUN go build -o /go/bin/hello
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello

############################
# STEP 2 build a small image
############################
FROM scratch
EXPOSE 8081
# Copy our static executable.

COPY --from=builder /go/bin/hello /go/bin/hello
# Run the hello binary.
ENTRYPOINT ["/go/bin/hello"]
# start from scratch
#FROM scratch
#
## copy our static linked executable
##COPY --from=builder /go/bin/hello ./hello
#COPY main ./hello
#
## tell how to run this container
#CMD ["./hello"]