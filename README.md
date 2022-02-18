# Linden Honey API

> Service with lyrics API powered by Go kit

[![build](https://img.shields.io/github/workflow/status/linden-honey/linden-honey-api-go/CI)](https://github.com/linden-honey/linden-honey-api-go/actions?query=workflow%3ACI)
[![version](https://img.shields.io/github/go-mod/go-version/linden-honey/linden-honey-api-go)](https://go.dev/)
[![coverage](https://img.shields.io/codecov/c/github/linden-honey/linden-honey-api-go)](https://codecov.io/github/linden-honey/linden-honey-api-go)
[![tag](https://img.shields.io/github/tag/linden-honey/linden-honey-api-go.svg)](https://github.com/linden-honey/linden-honey-api-go/tags)

## Technologies

- [Golang](https://go.dev/)
- [Go kit](https://gokit.io/)
- [MongoDB](https://www.mongodb.com/)

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

[https://linden-honey-api-go.herokuapp.com](https://linden-honey-api-go.herokuapp.com)
