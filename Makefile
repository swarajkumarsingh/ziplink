run:
	docker build -t ziplink . && docker run -p 8080:8080 ziplink

compose_build:
	docker compose build

compose:
	make compose_build && make only_compose

only_compose:
	docker compose up

build:
	docker build -t ziplink

start:
	docker run -p 8080:8080 ziplink

dev:
	nodemon --exec go run main.go

deploy: echo "TODO"

test: echo "TODO"