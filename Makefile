.PHONY: build
build: generate-idl
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
