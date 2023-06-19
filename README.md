# README #

## Setup

1. Install Go version 1.19
2. Install Mockery version v2.13 or later
3. Use GoLand (recommended)
4. Download dependencies with command `go mod download`
5. Create `.env` file based on `.env.example`

## Run

Use this command to run API app from root directory:

```shell
go run cmd/api/main.go
```

Use this command to run consumer app from root directory:

```shell
go run cmd/consumer/main.go
```

Use this command to run scheduler app from root directory:

```shell
go run cmd/scheduler/main.go
```

## Unit Tests

### Generate Mocks

To generate mock, run:

```shell
mockery --all --keeptree --case underscore --with-expecter --output ./test/mocks
```

### Run Unit Tests

To run unit tests:
```shell
go test ./...
```

---

## TODO README

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Summary of set up
* Configuration
* Dependencies
* Database configuration
* How to run tests
* Deployment instructions

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact# base-project
