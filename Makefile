.PHONY: compile
compile: generate-idl
	go build ./...

.PHONY: test
test: compile
	go test ./...

.PHONY: generate-idl
generate-idl: 
	./scripts/generate_protos.sh

.PHONY: build-python-idl
build-python-idl:
	@python setup.py sdist
