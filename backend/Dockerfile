FROM golang:latest AS builder

RUN go version

COPY . /github.com/poncheska/terminal-http-chat/backend/
WORKDIR /github.com/poncheska/terminal-http-chat/backend/

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o ./.bin/api ./cmd/api/main.go

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /github.com/poncheska/terminal-http-chat/backend/.bin/api .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

EXPOSE 8080

CMD /wait && ./api
