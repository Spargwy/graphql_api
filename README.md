Простой веб-сервис на GraphQL с использованием следующих библиотек:

- go-pg - для работы с PostgreSQL
- dbMate - для управления миграциями в проекте
- twilio/twilio-go - сервис и библиотека для взаимодействия с ним. Используется для отправки смс-кодов в процессе аутентификации пользователя.

# Настройка и запуск:
## Настройка:

```
git clone https://github.io/Spargwy/graphql_api
cd graphql_api
```
- Установка make

- В случае запуска через докер, единственное, что понадобится - это docker и docker-ccompose: 

https://docs.docker.com/engine/install/, 

https://docs.docker.com/compose/install/.

Для запуска без докера требуется установка следующих библиотек:
- Go: https://go.dev/
- gqlgen: https://gqlgen.com/
- PostgreSQL: https://www.postgresql.org/
- dbMate: https://github.com/amacneil/dbmate

А также установка зависимостей проекта командой
```go mod download```


В корневой директории представлен файл .env.example. По его подобию необходимо создать .env файл и заменить некоторые параметры вашими собственными. В случае с параметрами, связанными с twilio, требуется создать и настроить аккаунт: https://www.twilio.com/.

Пункт выше выполняется независимо от типа запуска.

## Запуск
Запуск с помощью докера:

- ```make docker-run``` - создание и запуск веб-сервиса с БД.
- ```make docker-migrate``` - накат миграций в БД внутри контейнера.
- Для заливки тестовых данных выполняется команда 
```
make docker-test-data
```

Без докера:

- Создать базу и накатить миграции:
```make migrate ```
- Залить тестовые данные:
```make local-test-data```
- Запустить сервер командами:
```
 go build -o server
 ./server
 ```

 Запуск линтера осуществляется с помощью golangci-lint через докер командой
 ```make lint``` или ```golangci-lint run ./...``` если имеется установленный golangci-lint: https://golangci-lint.run/usage/install/
