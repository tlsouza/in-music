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

clean:
	rm main

setup:
	go mod tidy

setup-dev: setup
	npm install -g nodemon

run-local:
	go build ./cmd/main.go
	./main

watch-local:
	nodemon --exec go run cmd/main.go --signal SIGTERM

docker-build:
	docker build -f ./dockerfile --build-arg node_env=${NODE_ENV} --build-arg app_version=${TAG} --build-arg app_name=${APP_NAME} --build-arg aws_region=${AWS_REGION} -t $(FINAL_URL):$(TAG) .
	docker build -f ./dockerfile --build-arg node_env=${NODE_ENV} --build-arg app_name=${APP_NAME} --build-arg aws_region=${AWS_REGION} -t $(FINAL_URL):latest .

run-docker: docker-build
	docker run -dp 3000:3000 $(FINAL_URL):latest

resources-local:
	docker-compose build
	docker-compose up -d

resources-redis-local:
	docker-compose build redis
	docker-compose up -d redis

unit-test:
	ENV=testing NEW_RELIC_ENABLED=false go test ./... -cover -v

integration-test:
	export export NEW_RELIC_ENABLED=false && go test ./app/ports/... -cover -v

test: unit-test integration-test

help:
	@echo "Commands:"
	@echo "	check:			Check golang, npm, docker and docker-compose versions"
	@echo "	setup:			Install npm modules(with yarn) to run application"
	@echo "	run-local: 		Build and run the app"
	@echo "	watch-local:		Build and run the app with nodemon"
	@echo "	docker-build:		Build docker image with dockerfile"
	@echo "	run-docker:		Build and run the app with docker"
	@echo "	test:			Run integration and unit tests"
	@echo "	unit-test:		Run unit tests only"
	@echo "	integration-test:	Run integration tests only"
	@echo "	load-test:		Run k6 tests"

check:
	@echo "\n*** Checking versions ***"
	@echo Go: $(shell go version)
	@echo npm: $(shell npm --version)
	@echo Docker: $(shell docker --version) 
	@echo Docker-Compose: $(shell docker-compose --version)
	@echo "\n"

### Terraform local scripts (dont forget to use ./env/.env file to setup vars)
tf-setup:
	make -C terraform setup
tf-init-plan:
	make -C terraform init
tf-plan:
	make -C terraform terraform-plan
tf-refresh:
	make -C terraform terraform-apply-refresh
tf-destroy:
	make -C terraform terraform-destroy
tf-destroy-single:
	make -C terraform terraform-destroy-single
tf-fmt:
	make -C terraform tf-fmt

### iamlive

# Setup: export AWS_CA_BUNDLE=~/.iamlive/ca.pem && export HTTP_PROXY=http://127.0.0.1:10080 && export HTTPS_PROXY=http://127.0.0.1:10080
# Clean: unset HTTP_PROXY && unset HTTPS_PROXY && unset AWS_CA_BUNDLE

iamlive-local:
	docker run --rm -v ~/.iamlive:/home/appuser/.iamlive -v ${PWD}:/app/output -p 10080:10080 --name iamlive -it vivareal/iamlive:latest --profile preprod

install-k6-mac:
	brew install k6
	
load-test:
	k6 run ./test/load/script.js -e ADYEN_USER=${ADYEN_USER} -e ADYEN_PWD=${ADYEN_PWD}

#Extract cops non-secret envs and generate .env
create-dot-env:
	./scripts/get-envs.sh ${COPS_ENV} ${COPS_UUID}

#Kafka-topics
gt:  ## Generate kafka-topics files
	@echo "Generating topics with name: ${N} and description: ${D} and resilience: ${R}"
	./scripts/generate-topics.sh "${N}" "${D}" "${R}"

gt-help: ## Help to generate topics
	@echo "Generate kafka-topics files using N=TOPIC_NAME and D="description text" and R(Y/n)=n #flag for retry and dlq topics"
	@echo "Example w/ retry & dlq:\n\t make gt N=topic_name D=\"description text\""
	@echo "Example w/o retry & dlq:\n\t make gt N=topic_name D=\"description text\" R=n"
	@echo "Output:\n./topics/preprod\n./topics/prod"

patch-cops:
	./scripts/restart-app-cops.sh ${COPS_NAMESPACE} ${APP_NAME}

update-repos:
	export GH_TOKEN=${GH_TOKEN} && ./scripts/update-repos.sh ${APP_NAME}
