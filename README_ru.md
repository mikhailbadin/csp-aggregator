# CSP Aggregator

CSP-aggregator это сервис, который принимает и обрабатывает [Content-Security-Policy Report](https://www.w3.org/TR/CSP2/). После приема CSP Report, логгирует его в [MongoDB](https://www.mongodb.com/), а также делает обработку, после которой отправляет его в [CSP-Store](https://github.com/mikhailbadin/csp-store) (который работает на [TarantoolDB](https://tarantool.io/)) для последующего анализа.

# Начало

Для запуска сервиса необходимо иметь:

1. [CSP-Aggregator](https://github.com/mikhailbadin/csp-aggregator)
2. [CSP-Store](https://github.com/mikhailbadin/csp-store)
3. [MongoDB](https://www.mongodb.com/)

# Установка

## C использованием Docker

### Требования

- [Docker](https://www.docker.com/)

### Сборка

Для сборки перейдите в корень проекта из выполните сборку из Dockerfile.

Пример:

```
docker build -t csp-aggregator:scratch .
```

## C использованием golang

### Зависимости

- [golang](https://golang.org/)
- [gin-gonic](github.com/gin-gonic/gin)
- [mgo](github.com/globalsign/mgo)
- [gotoenv](github.com/joho/godotenv)
- [go-tarantool](github.com/tarantool/go-tarantool)

### Сборка

Для установки сервиса с помощью `go get` введите в терминале:

```shell
go get -u github.com/mikhailbadin/csp-aggregator
```

После ввода этой команды приложение будет загружено и установлено в папку `$GOPATH/bin/`

Для сборки локально в папке проекта введите:

```shell
make go-compile
```

Скомпилированное приложение будет находится в папке `./bin`.

# Запуск

Приложение при запуске берет параметры из переменных окружения. Так же параметры могут быть описане в файле `.env`

Поддерживаются следующие параметры:

Конфигурация сервера:

- `SERVER_ADDR` - для указания порта, на котором сервер будет работать.

Конфигурация MongoDB:

- `MONGO_URI` - URI для подключения к MongoDB

Конфигурация для подлючение к [CSP-Store](https://github.com/mikhailbadin/csp-store) (TarantoolDB):

- `TARANTOOL_URL` - URL для подключения
- `TARANTOOL_USER` - имя пользователя
- `TARANTOOL_PASS` - пароль

Пример конфигурации:

```shell
# Server configuration
SERVER_ADDR=":8080"

# MongoDB configuration
MONGO_URI="127.0.0.1:27017"

# TarantoolDB configuration
TARANTOOL_URL="127.0.0.1:3301"
TARANTOOL_USER="guest"
TARANTOOL_PASS=""
```

# Работа с сервисом

Сервис имеет следующий API:

- `/csp_report` - для приема отчетов заголовка `Content-Security-Policy`
- `/csp_report_only` - для приема отчетов заголовка `Content-Security-Policy-Report-Only`
