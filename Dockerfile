FROM golang:1.14-alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /cameo
ENV CGO_ENABLED=0

COPY go.mod go.mod
RUN go mod download

COPY . .
RUN go test ./...
RUN GOOS=linux go build -o /cameo/cameo

FROM scratch
COPY --from=builder /cameo/cameo /
COPY etc/cameo /etc/cameo

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

VOLUME /etc/cameo
EXPOSE 8080

ENTRYPOINT ["/cameo"]
