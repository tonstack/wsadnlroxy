# -------- download modules
FROM golang:1.18-alpine3.16 AS modules

COPY go.mod go.sum /modules/
WORKDIR /modules

RUN go mod download

# -------- build the app
FROM golang:1.18-alpine3.16 as builder

COPY --from=modules /go/pkg /go/pkg

COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o wstcproxy ./cmd/main.go

# -------- run the app
FROM alpine:3.16 as app
COPY --from=builder /build/wstcproxy .
CMD ["./wstcproxy"]
