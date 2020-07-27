FROM golang:1.14-alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /cameo

COPY . .
RUN go mod download
RUN make test compile

FROM scratch
COPY --from=builder /cameo/cameo /
COPY etc/cameo /etc/cameo
COPY views /etc/cameo/views

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

VOLUME /etc/cameo
EXPOSE 8080

ENTRYPOINT ["/cameo"]
