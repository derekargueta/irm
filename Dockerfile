FROM golang:1.14-alpine

WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="./domains/finaldomaintest.txt" -w="30" -password="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"
