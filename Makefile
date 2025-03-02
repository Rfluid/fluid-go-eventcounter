EXCHANGE=eventcountertest
AMQP_PORT=5672
AMQP_UI_PORT=15672
CURRENT_DIR=$(shell pwd)
SLEEP_TIME=10

env-up:
	docker run -d --name evencountertest-rabbitmq -p $(AMQP_UI_PORT):15672 -p $(AMQP_PORT):5672 rabbitmq:3-management

env-down:
	docker rm -f evencountertest-rabbitmq

build-generator:
	go build -o bin/generator cmd/generator/*.go

generator-publish: build-generator
	bin/generator -publish=true -size=100 -amqp-url="amqp://guest:guest@localhost:$(AMQP_PORT)" -amqp-exchange="$(EXCHANGE)" -amqp-declare-queue=true

generator-publish-with-resume: build-generator
	bin/generator -publish=true -size=100 -amqp-url="amqp://guest:guest@localhost:$(AMQP_PORT)" -amqp-exchange="$(EXCHANGE)" -amqp-declare-queue=true -count=true -count-out="$(CURRENT_DIR)/data" -count=true

up: build-generator env-up

run:
	go run cmd/consumer/main.go -amqp-url="amqp://guest:guest@localhost:$(AMQP_PORT)" -amqp-exchange="$(EXCHANGE)"

# Acronym for Do Not Waste My Time
dnwmt: env-up
	sleep $(SLEEP_TIME) # Wait for SLEEP_TIME seconds to ensure RabbitMQ is ready
	$(MAKE) generator-publish
	$(MAKE) run
	$(MAKE) env-down
