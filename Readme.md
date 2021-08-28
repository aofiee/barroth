# Readme

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=alert_status)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=coverage)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=ncloc)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=security_rating)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=sqale_index)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=code_smells)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=aofiee_golang-clean-architecture&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=aofiee_golang-clean-architecture)

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

./node_modules/.bin/eslint --init