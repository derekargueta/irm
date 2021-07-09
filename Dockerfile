FROM golang:1.14-alpine

# Can be overridden at runtime using `docker run -e WORKERS=n ...`
ENV WORKERS=25

WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="/app/domains/finaldomaintest.txt" -w="$WORKERS"
