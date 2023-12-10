FROM golang:1.21-alpine

WORKDIR /app
COPY ./app /app

RUN apk update

RUN go mod init main \
    && go mod tidy \
    && go build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
EXPOSE 8080


CMD ["go", "run", "main.go"]