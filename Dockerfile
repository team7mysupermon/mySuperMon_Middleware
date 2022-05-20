# syntax=docker/dockerfile:1
FROM golang:latest
WORKDIR /src
COPY ./ ./

RUN go mod tidy
RUN go mod download
RUN go build -o /mysupermon-middleware-prometheus

EXPOSE 8999
EXPOSE 9090

CMD [ "/mysupermon-middleware-prometheus" ]
