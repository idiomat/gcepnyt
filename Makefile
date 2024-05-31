default: help

help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

e0: ## run e0
	go run ./e0-sequential/*.go

.PHONY: e0-with-profiling
e0-with-profiling: ## run e0-sequential with pprof
	go run ./e0-sequential-pprof/*.go

.PHONY: e0-view-profile
e0-view-profile: ## view e0-sequential goroutine profile
	go tool pprof -http=:6060 e0.goroutine.prof

.PHONY: e1
e1: ## run e1
	go run ./e1-synchronized/*.go

.PHONY: e1-with-profiling
e1-with-profiling: ## run e1-synchronized with pprof
	go run ./e1-synchronized-pprof/*.go

.PHONY: e1-view-profile
e1-view-profile: ## view e1-synchronized goroutine profile
	go tool pprof -http=:6060 e1-synchronized.goroutine.prof

e1-limited: ## run e1 with a ulimit
	ulimit -n 100; go run ./e1-synchronized/*.go -from=8000 -to=9000
