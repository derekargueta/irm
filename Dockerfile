FROM golang:1.14-alpine 

WORKDIR /app
COPY . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="/app/irm/domains/domaintest.txt"
