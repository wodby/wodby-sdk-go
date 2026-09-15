-include .env version

SWAGGER_CODEGEN_VER = 2.3.1
SWAGGER_CODEGEN_URL = https://repo.maven.apache.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_CODEGEN_VER)/swagger-codegen-cli-$(SWAGGER_CODEGEN_VER).jar
SWAGGER_CODEGEN_JAVA_OPTS = -Xmx1024M -DapiTests=false -DmodelTests=false
JAVA_IMAGE ?= eclipse-temurin:17-jre
UID ?= $(shell id -u)

default: build

test:
	go test ./...
.PHONY: test

build: clean codegen
.PHONY: build

update-readme:
	gotpl ./tpl/readme.tpl.md > README.md
.PHONY: update-readme

codegen:
	[ -f ./codegen.jar ] || wget -nv "$(SWAGGER_CODEGEN_URL)" -O ./codegen.jar
	docker run --rm --user "$(UID)" \
		-v "$(PWD)":/gen \
		-w /gen \
		"$(JAVA_IMAGE)" java $(SWAGGER_CODEGEN_JAVA_OPTS) -jar ./codegen.jar generate \
			-i ./swagger.json \
			-l go \
			-o ./pkg \
			-D packageName=client
	rm -f ./pkg/.travis.yml \
		./pkg/git_push.sh \
		./pkg/.gitignore
	./fix-generated-go.sh
.PHONY: codegen

clean:
	mv ./pkg/.swagger-codegen-ignore ./
	rm -rf ./codegen.jar ./pkg/*
	mv ./.swagger-codegen-ignore ./pkg/
.PHONY: clean
