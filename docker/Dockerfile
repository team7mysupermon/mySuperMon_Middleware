# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /src
COPY ./src ./

RUN go mod tidy
RUN go mod download
RUN go build -o /mysupermon-middleware-prometheus

EXPOSE 8081

CMD [ "/mysupermon-middleware-prometheus" ]
