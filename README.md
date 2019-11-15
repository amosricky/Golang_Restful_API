# Golang_Restful_API
A simple restful API.

## Installation
### Git clone from github
```
$ git clone https://github.com/amosricky/Golang_Restful_API.git
```

### Go get from github
```
$ go get github.com/amosricky/Golang_Restful_API
```

## How to run

### Required

- Golang >= 1.13
- Docker

### Run
```
$ cd $GOPATH/src/Golang_Restful_API

$ go run main.go 
```

### Run in container
```
$ cd $GOPATH/src/Golang_Restful_API

$ docker build -t golang-api . 

$ docker run -p 8000:8000 golang-api
```

## Features

- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable