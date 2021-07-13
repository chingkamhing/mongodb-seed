.PHONY: help
help:
	@echo "Usage:"
	@echo "    build                Build this project locally"
	@echo "    docker-up            docker-compose up"
	@echo "    docker-down          docker-compose down"

#
# Makefile commands
#

.PHONY: build
build:
	go build -o mongodb-seed .

#
# Docker commands
#

.PHONY: docker-up
docker-up:
	docker-compose -f docker-compose.yml up --build

# docker-compose down
.PHONY: docker-down
docker-down:
	docker-compose -f docker-compose.yml down
