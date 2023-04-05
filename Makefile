include .env
export $(shell sed 's/=.*//' .env)

build:
	go build ./cmd/stealcamoor
PHONY: build

generate:
	abigen --abi assets/Stealcam.json --pkg abis --type Stealcam --out pkg/abis/stealcam.go
PHONY: generate

run:
	./stealcamoor
.PHONY: run
