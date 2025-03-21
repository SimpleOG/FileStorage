# Используем официальный образ Go
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum   ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o FileStorage ./cmd

# Используем минимальный образ для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранный бинарник из builder-контейнера
COPY --from=builder /app/FileStorage .

# Указываем команду для запуска приложения
CMD ["./FileStorage"]