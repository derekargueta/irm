FROM golang:1.14-alpine

ENV pass=""

WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="./domains/finaldomaintest.txt" -w="30" -git=0
