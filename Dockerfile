FROM node:16-alpine as webui

WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY webui-vue/package.json /app/package.json

# install project dependencies
RUN yarn install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY webui-vue/. /app/.

# build app for production with minification
RUN yarn run build
 
FROM golang:1.16.3-alpine as release

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/lusha/go-template-online

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY main.go .
COPY go.mod .

# Download all the dependencies
RUN go get -v
RUN go build

COPY --from=webui /app/dist webui-vue/dist 

# This container exposes port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD ["gotplonline"]
