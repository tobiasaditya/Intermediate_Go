FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod tidy

COPY . ./

RUN go build -o main

EXPOSE 8000

CMD ["./main"]