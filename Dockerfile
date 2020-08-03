FROM golang:1.14-alpine
WORKDIR /app
COPY . /app
RUN go build

FROM golang:1.14-alpine
WORKDIR /app
COPY --from=0 /app/httpproxy-go-runner .
ENTRYPOINT ["/app/httpproxy-go-runner"]
