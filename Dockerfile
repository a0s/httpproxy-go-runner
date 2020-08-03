FROM golang:1.14-alpine AS builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/httpproxy-go-runner

FROM scratch
WORKDIR /app
COPY --from=builder /app/httpproxy-go-runner .
ENTRYPOINT ["/app/httpproxy-go-runner"]
