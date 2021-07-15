FROM golang:1.14-alpine 

ENV WORKERS=$(nproc)

WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT go run ./cmd/analyze -f="/app/domains/finaldomaintest.txt" -o="/app/cmd/analyze/results/results.csv" -w="$WORKERS" 