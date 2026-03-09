FROM golang:1.22 AS builder
WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/server /app/server

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
