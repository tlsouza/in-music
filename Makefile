TAG=latest
IMAGE = fp-template-go-api
CONTAINER_REGISTRY_HOST=073521391622.dkr.ecr.us-east-1.amazonaws.com
REPOSITORY=cross/fp
FINAL_URL=$(CONTAINER_REGISTRY_HOST)/${REPOSITORY}/$(IMAGE)
APP_NAME=
NODE_ENV=
AWS_REGION=us-east-1
ADYEN_USER=test
ADYEN_PWD=password
COPS_UUID=
COPS_ENV=


setup:
	go mod tidy

setup-dev: setup
	npm install -g nodemon

run-local:
	go build ./cmd/main.go
	./main

unit-test:
	ENV=testing NEW_RELIC_ENABLED=false go test ./... -cover -v

integration-test:
	export export NEW_RELIC_ENABLED=false && go test ./app/... -cover -v

test: unit-test integration-test