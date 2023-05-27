FROM golang:1.19-buster as builder

WORKDIR /chat
COPY . .

WORKDIR /chat/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ../chat

FROM alpine:3.15.4
WORKDIR /chat
COPY --from=builder /chat/chat .
CMD ["./chat"]
