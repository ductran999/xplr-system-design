.PHONY: proto
proto: ## Generate Go code by buf
	@echo "Linting and generating code with buf..."
	buf lint
	buf generate --clean
	@echo "Done!"

.PHONY: server
server: ## Run server
	go run server/main.go

.PHONY: client
client: ## Run client server
	go run client/main.go

.PHONY: help
help: ## Show list of available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

