FROM golang:1.19-buster as builder

WORKDIR /auth
COPY . .

WORKDIR /auth/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../auth

FROM alpine:3.15.4
WORKDIR /auth
COPY --from=builder /auth/auth .
CMD ["./auth"]
