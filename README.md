# Быстрый поиск анаграмм в словаре

## Установка

- `go get`
- `go run main.go`

## Описание эндпоинтов

### Получение версии приложения
```
curl --request GET \
  --url http://localhost:8080/
```

### Загрузка списка анаграмм
```
curl --request POST \
  --url http://localhost:8080/load \
  --header 'content-type: application/json' \
  --data '["foobar", "aabb", "baba", "boofar", "test", "ttse"]'
```

### Получение анаграмм слова
```
curl --request GET \
  --url 'http://localhost:8080/get?word=aabb'
```