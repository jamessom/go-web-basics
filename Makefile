# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

%:
	@:

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

# DOCKER TASKS

build: ## Build the containers
	docker compose build

migrate: ## Create tables
	docker compose --profile tools run go-migrate.gowebbasics.com

up: ## Start all containers
	docker compose up -d

down: ## Stop all containers
	docker compose down

log: ## Show logs from containers. Only app and db
	docker compose logs -f go-$(call args,defaultstring).gowebbasics.com