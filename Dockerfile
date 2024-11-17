# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Устанавливаем необходимые пакеты
RUN apk update && apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка приложения
RUN go build -o web-chat-backend main.go

# Stage 2: Создаем минимальный образ для выполнения приложения
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=builder /app/web-chat-backend .

# Открываем порт 8080
EXPOSE 8080

# Команда для запуска приложения
CMD ["./web-chat-backend"]
