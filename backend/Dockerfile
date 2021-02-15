FROM golang:latest AS builder

RUN go version

COPY . ./github.com/poncheska/terminal-http-chat/backend/
WORKDIR /github.com/poncheska/terminal-http-chat/backend/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest

COPY ./wait-for-postgres.sh .
COPY --from=0 /github.com/poncheska/terminal-http-chat/.bin/app .

CMD ["./app"]
