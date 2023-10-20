##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	go test -v -vet=off -count=1 -timeout 5m -v ./...

##@ Build

.PHONY: build
build: fmt vet ## Build sonic binary.
	go build -a -trimpath -o bin/sonic cmd/main.go
