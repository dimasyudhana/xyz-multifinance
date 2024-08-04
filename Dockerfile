FROM golang:1.20-alpine

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o backend-service

CMD ["./backend-service"]