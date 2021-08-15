# Readme

[![Quality gate](https://sonarcloud.io/api/project_badges/quality_gate?project=aofiee_golang-clean-architecture)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)

## Setup

1. docker-compose up --build
2. docker-compose up -d
3. docker-compose logs -t --follow
4. go test ./...
   > or
5. go test -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html

## Database

> PHPMyAdmin http://localhost:8080/
> mariadb port 3308

## Redis

> port 6379
> password = password

## RabbitMq

> port 5672
> http://localhost:15672/
> user = guest , password = guest

## Mockery

go get github.com/vektra/mockery/.../
mockery -all -recursive -dir=./domains
