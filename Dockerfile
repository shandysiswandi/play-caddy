FROM golang:1.15-alpine as builder
WORKDIR /appgo
COPY . .
RUN go build

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /appgo/playcaddy /app/playcaddy
ENTRYPOINT ["/app/playcaddy"]