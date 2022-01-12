FROM golang:1.17-alpine
RUN CGO_ENABLED=0
RUN apk add build-base
WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT ["go", "test"]

