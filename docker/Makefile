API_BINARY=backend

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_backend
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	docker rmi docker-backend
	@echo "Done!"

## build_backend: builds the api binary as a linux executable
build_backend:
	@echo "Building api binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${API_BINARY} ../cmd/api
	@echo "Done!"