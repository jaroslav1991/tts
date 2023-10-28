# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем содержимое текущей директории (где находится Dockerfile) в контейнер
COPY . .

# Собираем Go приложение
RUN go build -o main cmd/cli/main.go

# Указываем команду, которая будет выполняться при запуске контейнера
CMD ["./main"]

#### Just comments for test

# Собираем образ
# docker build -t tts-cli .

# docker run --name tts-cli-1 -it --rm -p9292:9191 tts-cli
# docker run --name tts-cli-2 -it --rm -p9393:9191 tts-cli
# docker run --name tts-cli-3 -it --rm -p9494:9191 tts-cli
# docker run --name tts-cli -it --rm -v /Users/andrey.tikhonov/go/src/github.com/jaroslav1991/tts/outbox:/app/outbox tts-cli ./main -d '{"uid":"qwerty123","pluginType":"jetbrains","pluginVersion":"1.0.0","ideType":"intellij idea","ideVersion":"2.1.1","events":[{"createdAt":"2022-01-1114:23:01","type":"modifyfile","project":"someproject","projectBaseDir":"./","language":"golang","target":"C/Projects/Golang/cli-tts"}]}'
