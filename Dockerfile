FROM golang:1.16.3-alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/meteorops/go-template-online

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY main.go .
COPY go.mod .

# Download all the dependencies
RUN go get -v
RUN go build

COPY webui-vue/dist webui-vue/dist 

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["gotplonline"]
