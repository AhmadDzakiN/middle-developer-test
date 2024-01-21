FROM golang:1.20.13-alpine3.19

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

COPY params/.env .

RUN go mod tidy

RUN go build -o /middle-developer-test

EXPOSE 8000

CMD ["/middle-developer-test"]