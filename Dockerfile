# Используем golang:alpine для сборки
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Устанавливаем необходимые зависимости
RUN apk add --no-cache gcc musl-dev ca-certificates git

# Копируем файлы go.mod и go.sum, чтобы кешировать зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарного файла
RUN go build -o main .

# --- Финальный контейнер ---
FROM alpine:latest

WORKDIR /root/

# Устанавливаем ca-certificates (для работы HTTPS)
RUN apk add --no-cache ca-certificates

# Копируем скомпилированное приложение
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
