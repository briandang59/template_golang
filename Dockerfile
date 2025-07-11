# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o be_scada ./cmd/api

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/be_scada .
COPY .env .

EXPOSE 5000
CMD ["./be_scada"]
