# Используем официальный образ Go как основу
FROM golang:alpine

# Устанавливаем рабочий каталог в /app
WORKDIR /app

# Копируем исходный код в рабочий каталог
COPY . /app

# Устанавливаем переменные среды
ENV GOOS=linux
ENV GOARCH=amd64

# Сборка проекта
RUN go build -o main cmd/main.go

# Установка порта для запуска приложения
EXPOSE 8080

# Запуск приложения
CMD ["./main"]