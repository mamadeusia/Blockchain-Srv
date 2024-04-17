GOPATH:=$(shell go env GOPATH)
include .env

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/BlockchainSrv.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o BlockchainSrv *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t BlockchainSrv:latest .

.PHONY: ready
ready:
	@make update
	@make tidy
	@make proto

.PHONY: run
run:
	@make build
	@./BlockchainSrv

.PHONY: migration-update
migration-update:
	@liquibase update --changelog-file=database/changelog.yaml --url=jdbc:postgresql://localhost:5432/blockchainsrv --username=postgres --password=postgres

.PHONY: migration-rollback
migration-rollback:
	@liquibase rollback-count 1 --changelog-file=database/changelog.yaml --url=jdbc:postgresql://localhost:5432/blockchainsrv --username=postgres --password=postgres