# Alpine is awesome and lightweight
FROM golang:1.8.3-alpine
# Install some things to make development easier
RUN apk update && apk add git bash mysql-client
# Set up a go environment
RUN mkdir -p /app/src/github.com/domtheporcupine/updoots
RUN mkdir -p /app/bin
RUN mkdir -p /app/pkg
# Set go environment variables
ENV GOPATH /app
ENV PATH $PATH:/app/bin
# Install our go dependencies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gorilla/mux
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/codegangsta/gin

# Build our app
WORKDIR /app/src/github.com/domtheporcupine/updoots

# Start the API, using gin for auto reload/compile
CMD ["gin", "run", "app"]

