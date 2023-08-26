FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app

COPY . .
RUN go build -o main main.go
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.17 AS stage

WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate /app/migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migrations ./db/migrations

EXPOSE 8000
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]