-include .env version

OPENAPI_GENERATOR_VER = v7.10.0
OPENAPI_GENERATOR_IMAGE = openapitools/openapi-generator-cli:$(OPENAPI_GENERATOR_VER)
UID ?= $(shell id -u)

default: build

build: clean codegen
.PHONY: build

update-readme:
	gotpl ./tpl/readme.tpl.md > README.md
.PHONY: update-readme

codegen:
	docker run --rm \
		-v "$(PWD)":/gen \
		-w /gen \
		"$(OPENAPI_GENERATOR_IMAGE)" generate \
			-i ./swagger.json \
			-g go \
			-o ./pkg \
			--additional-properties=packageName=client
	sudo chown -R $(UID) ./
	rm -f ./pkg/.travis.yml \
		./pkg/git_push.sh \
		./pkg/.gitignore
.PHONY: codegen

clean:
	mv ./pkg/.openapi-generator-ignore ./
	rm -rf ./codegen.jar ./pkg/*
	mv ./.openapi-generator-ignore ./pkg/
.PHONY: clean
