FROM golang:latest AS builder

RUN go version
ENV GOPATH=/

COPY ./ ./github.com/poncheska/terminal-http-chat/
WORKDIR ./github.com/poncheska/terminal-http-chat/

RUN go mod download

RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder ./github.com/poncheska/terminal-http-chat/.bin/app .

CMD ["./app"]
