install:
	npm i

dev:
	air || npm run dev || go run cmd/main.go

start: install dev

docker up:
	docker compose up

docker down:
	docker compose down

build:
	go build -o bin/main cmd/main.go

run-build: build
	bin/main
