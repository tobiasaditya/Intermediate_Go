# FROM golang:1.18-alpine

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download
# RUN go mod tidy

# COPY . ./

# RUN go build -o main

# EXPOSE 8000

# CMD ["./main"]

FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .

FROM alpine:3.13
WORKDIR /app
RUN pwd
COPY --from=builder /app/deploy/session-login-9 .
COPY --from=builder /app/template /app/template/.
RUN ls
EXPOSE 8000

ENV APP_NAME=session-login-9
CMD ["./session-login-9"]