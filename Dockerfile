FROM golang:latest

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go install github.com/cosmtrek/air@latest