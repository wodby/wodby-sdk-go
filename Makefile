-include .env version

SWAGGER_CODEGEN_VER = 2.3.1
SWAGGER_CODEGEN_URL = http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_CODEGEN_VER)/swagger-codegen-cli-$(SWAGGER_CODEGEN_VER).jar
SWAGGER_CODEGEN_JAVA_OPTS = -Xmx1024M -DapiTests=false -DmodelTests=false
MAVEN_VER = 3-jdk-7-alpine
UID ?= $(shell id -u)
API_BRANCH = $(shell cat ./.wodby-api/BRANCH)
API_VERSION = $(shell cat ./.wodby-api/VERSION)

default: none

none:
	@echo 'Target not specified'
.PHONY: none

export API_BRANCH
export API_VERSION
update-readme:
	gotpl ./tpl/readme.tpl.md > README.md
.PHONY: update-readme

codegen:
	[ -f ./codegen.jar ] || wget -nv "$(SWAGGER_CODEGEN_URL)" -O ./codegen.jar
	docker run -it --rm \
		-v "$(PWD)":/gen \
		-w /gen \
		maven:"$(MAVEN_VER)" java $(SWAGGER_CODEGEN_JAVA_OPTS) -jar ./codegen.jar generate \
			-i ./swagger.json \
			-l go \
			-o ./pkg \
			-D packageName=wodby-api
	sudo chown -R $(UID) ./
	rm -f ./pkg/.travis.yml \
		./pkg/git_push.sh \
		./pkg/.gitignore
.PHONY: codegen