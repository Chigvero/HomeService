FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Собираем приложение
RUN go build -o main cmd/app/main.go

# Указываем команду для запуска приложения
CMD ["./main"]

# Открываем порт 8081
EXPOSE 8080