FROM golang:1.16.4-alpine3.13 as builder
COPY go.mod go.sum /go/src/