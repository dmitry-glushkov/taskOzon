## Запуск

```sh
docker-compose up
```
---

## Запросы

### Создание новой записи
Принимает значения *key* (string), *value* (строки, списки, числа, мапы и тд.) и, опционально, *ttl* (int) - время жизни записи  в секундах:
```sh
curl --header "Content-Type: application/json" \
--request POST \
--data '{"key": "<KEY>", "value": <VALUE>, "ttl": 3}' \
http://localhost:8080/set
```

### Получение записи по ключу
Принимает значение ключа *key*:
```sh
curl --header "Content-Type: application/json" \
--request POST \
--data '{"key": "<KEY>"}' \
http://localhost:8080/get
```

### Получение всех существующих ключей
```sh
curl --header "Content-Type: application/json" \
--request GET \
http://localhost:8080/getallkeys
```

### Удаление записи
Удаляет запись по значению ключа *key*
```sh
curl --header "Content-Type: application/json" \
--request POST \
--data '{"key": "<KEY>"}' \
http://localhost:8080/del
```
---
## Benchmark
Результаты теста производительности для метода *set*
| Всего операций |  ns/op  |  B/op  | Allocs/op |
| -------------- | ------- | ------ | --------- |
|     53947      |  22058  |  3844  |    49     |
