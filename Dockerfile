FROM golang:1.19-alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -extldflags '-static'"

FROM alpine

COPY --from=builder /app/hello-hostname /app/hello-hostname
EXPOSE 5000

ENTRYPOINT ["/app/hello-hostname"]
