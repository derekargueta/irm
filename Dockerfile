        FROM golang:1.17-alpine

        ENV pass=""

        WORKDIR /app
        COPY . .
        ADD . .
        RUN apk --no-cache add curl
        RUN go get -d -v ./...
        ENTRYPOINT go run ./cmd/analyze -f="./domains/testset.txt" -w="30" -git=1

        # url: go run ./cmd/analyze -url="" -w="30" -git=0
        # file: go run ./cmd/analyze -f="./domains/testset.txt" -w="30" -git=0
