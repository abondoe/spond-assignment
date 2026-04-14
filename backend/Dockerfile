# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Kopier go.mod og go.sum først for å utnytte caching av dependencies
COPY go.mod go.sum ./
RUN go mod download

# Kopier resten av koden
COPY . .

# Bygg appen
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage 2: Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Kopier binærfilen fra builder-stadiet
COPY --from=builder /app/main .

# Eksponer porten appen kjører på
EXPOSE 8080

CMD ["./main"]