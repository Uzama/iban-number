FROM golang:1.18.2-alpine3.16

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o iban-number

EXPOSE 8080

ENTRYPOINT ["./iban-number"]