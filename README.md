# MalDB
MalDB - зеркало MyAnimeList на Golang

## Запуск в Docker
На данный момент сделан только локальный docker-compose (в нем запускается air для лайв ребилда кода). 
В будущем будет сделан и обычный docker-compose для прода, проксироваться будет через Traefik.

### Шаги запуска
- Для начала надо скопировать .env.example в .env файл

```bash
cd /_docker
cp .env.example .env
```

- После запускаем проект (дев среда, пока только так)
```bash
cd /_docker
docker-compose -f docker-compose-local.yml up --build
```


## Запуск без докера
- Копируем .env.example из _docker каталога в корень проекта

- Скачиваем и чистим зависимости
```bash
go mod download
go mod tidy
```

- билдим и запускаем проект (ну или делаем это через IDE)
```bash
go build -v -o ./bin/app ./cmd
./bin/app
```