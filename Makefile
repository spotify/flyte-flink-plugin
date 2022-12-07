.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: fmt
fmt: ## Run go fmt
	go fmt ./...

.PHONY: tidy
tidy: ## Run go mod tidy
	go mod tidy

.PHONY: build
build: generate-idl fmt vet tidy
	go build ./...

.PHONY: test
test: build
	go test ./...

.PHONY: generate-idl
generate-idl: 
	./scripts/generate_protos.sh

.PHONY: build-python-idl
build-python-idl:
	@python setup.py sdist
