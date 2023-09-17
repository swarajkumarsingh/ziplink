run:
	docker build -t ziplink . && docker run -p 8080:8080 ziplink

build:
	docker build -t ziplink

start:
	docker run -p 8080:8080 ziplink

dev:
	nodemon --exec go run main.go

deploy: echo "TODO"

test: echo "TODO"