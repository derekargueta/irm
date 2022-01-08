FROM golang:1.17-alpine

ENV pass=""

WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="./domains/testset.txt" -w="30" -git=0

# url: go run ./cmd/analyze -w="30" -git=0
# file: go run ./cmd/analyze -f="./domain 
# go run ./cmd/analyze -f="./domains/testset.txt" -w="30" -git=0