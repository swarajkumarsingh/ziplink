SCRIPT_FOLDER_NAME = commands

run:
	docker compose build
	docker compose up

build:
	docker build -t ziplink . && docker run -p 8080:8080 ziplink

start:
	docker run -p 8080:8080 ziplink

compose:
	docker compose build && docker compose up

down:
	docker compose down

dev:
	nodemon --exec go run main.go

install:
	go mod tidy

deploy: 
	echo "TODO"

test: 
	echo "TODO"

.PHONY: build run logs dockerstop
.SILENT: build run logs dockerstop