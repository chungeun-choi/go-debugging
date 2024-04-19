FROM golang:1.22

ENV CGO_ENABLED=0

RUN mkdir /app
WORKDIR /app
ADD . /app

RUN go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENTRYPOINT CompileDaemon --build="go build main.go" --command="dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec main"