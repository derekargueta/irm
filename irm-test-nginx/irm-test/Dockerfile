FROM golang:1.17-alpine
RUN CGO_ENABLED=0
RUN apk add build-base
COPY localhost.crt /usr/local/share/ca-certificates/

RUN cat /usr/local/share/ca-certificates/localhost.crt >> /etc/ssl/certs/ca-certificates.crt && \
    apk --no-cache add \
        curl
WORKDIR /app
COPY . .
ADD . .
RUN go get -d -v ./...
ENTRYPOINT ["go", "test"]

