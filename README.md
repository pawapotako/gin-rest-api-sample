## gin-rest-api-sample
Golang REST API sample with MySql integration using Gin and GORM.

This project is a sample project that contains following features:

- REST API server with [Gin Framework](https://github.com/gin-gonic/gin)
- Database integration using [GORM](http://gorm.io/)
- JWT Token based Authentication
- N Tier Architecture [N Tier Architecture Concept](https://stackify.com/n-tier-architecture/)

## Project Setup

```
$ go get github.com/jinzhu/gorm
$ go get github.com/spf13/viper
```

GORM should be installed via `go get`.

## MySql Database Setup

```
$ docker-compose up
```

MySql is implemented in Docker Compose. Docker should be installed first [Docker Desktop](https://www.docker.com/products/docker-desktop/).

## Start Project

```
$ go run main.go
```

