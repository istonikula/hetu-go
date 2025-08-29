.PHONY: help
help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: test
test: ## Run tests
	@go tool gotestsum --format=testname

.PHONY: lint
lint: ## Run lint
	@go tool golangci-lint run

.PHONY: check_upgrades
check_upgrades: ## List available minor/patch upgrades of required dependencies
	@# https://go.dev/wiki/Modules#how-to-upgrade-and-downgrade-dependencies
	@go list -u \
	-f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' \
	-m all 2> /dev/null
