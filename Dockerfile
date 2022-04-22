# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /little-mw-app

EXPOSE 8081

CMD [ "/little-mw-app" ]
