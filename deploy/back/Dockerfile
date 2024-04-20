FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY .. .

RUN go mod download

RUN go build -o main backend/cmd/backend/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

COPY /deploy /build/deploy

CMD ["./main", "--rootPath", "."]