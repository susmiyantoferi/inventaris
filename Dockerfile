FROM golang:1.22.2-bookworm AS builder
WORKDIR /app
COPY . .
COPY .env .
RUN go build -o goapp .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/goapp .
COPY --from=builder /app/.env .
CMD [ "./goapp" ]