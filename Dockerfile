FROM golang:1.22.2-bookworm AS builder
WORKDIR /app
COPY . .
COPY .env .
RUN go build -o goapp .

FROM debian:bookworm-slim
WORKDIR /app

# Install netcat (nc)
RUN apt-get update && apt-get install -y netcat && apt-get clean
COPY --from=builder /app/goapp .
COPY --from=builder /app/.env .
COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh

CMD ["./wait-for-db.sh", "./goapp" ]