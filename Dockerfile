# Alpine is awesome and lightweight
FROM golang:1.8.3-alpine
# Install some things to make development easier
RUN apk update && apk add git bash
# Set up a go environment
RUN mkdir -p /app/src/github.com/domtheporcupine/pittchat
RUN mkdir -p /app/bin
RUN mkdir -p /app/pkg
# Set go environment variables
ENV GOPATH /app
ENV PATH $PATH:/app/bin
# Install our go dependencies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/codegangsta/gin
# Add all the go files in our local directory into
# the container
ADD ./*.go /app/src/github.com/domtheporcupine/pittchat
# Build our app
RUN go install github.com/domtheporcupine/pittchat
WORKDIR /app/src/github.com/domtheporcupine/pittchat

# Start the API
CMD ["gin", "run", "project"]

