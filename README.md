# Linden Honey Go

> RESTful Web Service with lyrics powered by Golang

[![build](https://img.shields.io/github/workflow/status/linden-honey/linden-honey-go/CI)](https://github.com/linden-honey/linden-honey-go/actions?query=workflow%3ACI)
[![version](https://img.shields.io/github/go-mod/go-version/linden-honey/linden-honey-go)](https://golang.org/)
[![report](https://goreportcard.com/badge/github.com/linden-honey/linden-honey-go)](https://goreportcard.com/report/github.com/linden-honey/linden-honey-go)
[![coverage](https://img.shields.io/codecov/c/github/linden-honey/linden-honey-go)](https://codecov.io/github/linden-honey/linden-honey-go)
[![release](https://img.shields.io/github/release/linden-honey/linden-honey-go.svg)](https://github.com/linden-honey/linden-honey-go/releases)
[![reference](https://pkg.go.dev/badge/github.com/linden-honey/linden-honey-go.svg)](https://pkg.go.dev/github.com/linden-honey/linden-honey-go)

## Technologies

- [Golang](https://golang.org/)
- [Go kit](https://gokit.io/)

## Usage

### Local

Build application artifacts:

```bash
make build
```

Run tests:

```bash
make test
```

Run application:

```bash
make run
```

### Docker

Bootstrap full project using docker-compose:

```bash
docker-compose up
```

Bootstrap project excluding some services using docker-compose:

```bash
docker-compose up --scale [SERVICE=0...]
```

Stop and remove containers, networks, images:

```bash
docker-compose down
```

## Application instance

https://linden-honey-go.herokuapp.com
