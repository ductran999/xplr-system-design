.PHONY: proto
proto: ## Generate Go code by buf
	@echo "Linting and generating code with buf..."
	buf lint
	buf generate --clean
	@echo "Done!"

.PHONY: master
master: ## Run Master server
	cd services/master && go run main.go

slave: ## Run slave server
	cd services/slave && SERVER_URL=localhost:50051 REGISTRATION_TOKEN=OK VERSION=0.0.0 go run main.go

.PHONY: help
help: ## Show list of available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

