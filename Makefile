# include .env

BINARY=engine

engine:
	go build -o ${BINARY} app/*.go

test:
	go test -v ./...