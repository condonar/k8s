# ЭТАП 1: Сборка (Builder)
# Убрали --platform, так как на Intel это архитектура по умолчанию
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Копируем go.mod. Если его нет, команда 'go mod init' создаст его сама
COPY go.mod ./

# Копируем исходный код. Убедись, что файл называется index.go, раз ты его так вызываешь
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o server index.go

# ЭТАП 2: Финальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем результат сборки
COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]
 