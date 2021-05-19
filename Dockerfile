FROM golang:1.16.4-alpine3.13 as builder
COPY go.mod go.sum /go/src/github.com/kjuchniewicz/go-web-app/
WORKDIR /go/src/github.com/kjuchniewicz/go-web-app
RUN go mod download
COPY . /go/src/github.com/kjuchniewicz/go-web-app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-web-app ./cmd/web

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/kjuchniewicz/go-web-app /usr/bin/go-web-app
EXPOSE 8088 8088
ENTRYPOINT [ "/usr/bin/go-web-app" ]
