# CSP Aggregator

CSP aggregator is a service that accepts and processes [Content-Security-Policy Report](https://www.w3.org/TR/CSP2/). After receiving the CSP Report, logs it to [MongoDB](https://www.mongodb.com/), and also does the processing, after which it sends it to [CSP-Store](https://github.com/mikhailbadin/csp-store) (works in [TarantoolDB](https://tarantool.io/)) for further analysis.

# Introduction

To start the service, you must have:

1. [CSP-Aggregator](https://github.com/mikhailbadin/csp-aggregator)
2. [CSP-Store](https://github.com/mikhailbadin/csp-store)
3. [MongoDB](https://www.mongodb.com/)

# Building

## Using Docker

### Dependencies

- [Docker](https://www.docker.com/)

### How to build

Go to the root of the project and build from the Dockerfile.

Example:

```
docker build -t csp-aggregator:scratch .
```

## Using golang

### Dependencies

- [golang](https://golang.org/)
- [gin-gonic](github.com/gin-gonic/gin)
- [mgo](github.com/globalsign/mgo)
- [gotoenv](github.com/joho/godotenv)
- [go-tarantool](github.com/tarantool/go-tarantool)

### How to build

To install the service using `go get` enter in the terminal:

```shell
go get -u github.com/mikhailbadin/csp-aggregator
```

After entering this command, the application will be downloaded and installed in the folder: `$GOPATH/bin/`

To build locally in the project folder, enter:

```shell
make go-compile
```

The compiled application will be located in the folder `./bin`.

# How to run

The application takes parameters from environment variables at startup. Also, parameters can be described in the `.env` file.

The following parameters are supported:

Server configuration:

- `SERVER_ADDR` - to specify the port on which the server will work.

MongoDB configuration:

- `MONGO_URI` - URI to connect to MongoDB

[CSP-Store](https://github.com/mikhailbadin/csp-store) (TarantoolDB) configuration:

- `TARANTOOL_URL` - URI to connect to TarantoolDB
- `TARANTOOL_USER` - username
- `TARANTOOL_PASS` - password

Example configuration:

```shell
# Server configuration
SERVER_ADDR=":8080"

# MongoDB configuration
MONGO_URI="127.0.01:27017"

# TarantoolDB configuration
TARANTOOL_URL="127.0.0.1:3301"
TARANTOOL_USER="guest"
TARANTOOL_PASS=""
```

# Work with the service

The service has the following API:

- `/csp_report` - for receiving reports header `Content-Security-Policy`
- `/csp_report_only` - for receiving reports header `Content-Security-Policy-Report-Only`
