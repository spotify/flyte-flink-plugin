.PHONE: compile
compile: generate-idl
	go build ./...

.PHONY: generate-idl
generate-idl: 
	./scripts/generate_protos.sh

.PHONY: build-python-idl
build-python-idl:
	@python setup.py sdist
