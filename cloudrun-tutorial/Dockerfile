FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /app/main .

FROM alpine:3.14

ENV FUNCTION_TARGET=HelloWorld \
    LOCAL_ONLY=false

# ENV LOCAL_ONLY=true

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
