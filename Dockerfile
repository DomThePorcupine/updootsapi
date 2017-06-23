FROM golang:1.8.3-alpine
RUN apk update && apk add bash
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./app.go
CMD ["./app]

