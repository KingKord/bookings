SHELL=cmd.exe
BINARY=bookingApp

## up: start container in the background without forcing build
up:
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker started!"

## docker_up_build: Build all projects and start docker compose
up_build: build_bookings
	@echo "Starting docker images..."
	docker-compose up --build -d
	@echo "Docker started!"

build_bookings:
	@echo Building bookings binary...
	cd .\dockerfiles\ && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BINARY} ./../cmd/web
	@echo Done!

## docker_down: Stop docker compose
down:
	@echo "Stopping docker images..."
	docker-compose down
	@echo "Docker stopped!"

stop:
	@echo "Stopping app..."
	@-pkill -SIGTERM -f "./bookingApp"
	@echo "Stopped app!"

restart: down up_build up