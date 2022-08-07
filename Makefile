.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
		docker build -t harukiido/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
		docker compose build --no-cache

up:
		docker compose up -d

down:
		docker compose down

logs:
		docker compose logs -f

ps:
		docker compose ps

test:
		go test -race -shuffle-on ./...