.PHONY: help
help: ## Show list of available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: order-svc
order-svc: ## Run order server
	cd order-svc && go run main.go

.PHONY: notify-svc
notify-svc: ## Run user server
	cd notify-svc && go run main.go

setup: # setup infra
	docker-compose up -d

down: # cleanup infra
	docker-compose down -v