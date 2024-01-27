FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /app/main .

FROM alpine:3.14

ENV FUNCTION_TARGET=HelloWorld

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
