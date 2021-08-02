FROM golang:1.14-alpine

ENV pass=""

WORKDIR /app
COPY . .
COPY ./tempirmdata/irm-data/.git /tempirmdata/irm-data/.git
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="./domains/finaldomaintest.txt" -w="30" -password=""
