FROM golang:1.16-alpine3.13 as builder

WORKDIR /go/src/github.com/linden-honey/linden-honey-go

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY config ./config
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go install -v -ldflags="-w -s" ./cmd/server

FROM alpine:3.13

WORKDIR /app

ENV SERVER_ADDR=0.0.0.0:80

COPY --from=builder /go/bin/server /bin/server
COPY api/ ./api

ENTRYPOINT [ "/bin/server" ]
