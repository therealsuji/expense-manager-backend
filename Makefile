#!make
include .env

.PHONY: build
build:
	go build -o bin/app

.PHONY: dev
dev:
	# Custom port
	# air -- --port=:4002
	air

.PHONY: run
run: build
	./bin/app

.PHONY: test	
test:
	go test -v ./... -count=1

.PHONY: goose-status
goose-status:
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" status 

.PHONY: goose-create
goose-create:
ifdef NAME
	@echo "Creating migration $(NAME)"
	goose -dir ./migrations create $(NAME) sql
else
	$(error NAME is not set. Please provide a name using NAME=<migration-name>)
endif

.PHONY: goose-up
goose-up:
	@echo "Running migrations..."
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" up

.PHONY: goose-down
goose-down:
	@echo "Rolling back migrations..."
	goose -dir ./migrations postgres "user=$(DB_USERNAME) password=$(DB_PASSWORD) dbname=$(DB_DATABASE) sslmode=disable" down