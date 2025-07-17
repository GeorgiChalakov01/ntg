.PHONY: clear backend backend-logs database database-logs database-connect objectstorage objectstorage-logs all up down restart

clear:
	@clear

backend:
	@echo "=== Backend ==="
	@echo "Compiling source code..."
	@cd backend/app && make compile
	@echo "Building image..."
	@cd backend && docker buildx build -t ntg .
	@echo "Recreating container..."
	@docker compose up -d --no-deps --force-recreate backend

backend-logs:
	@docker logs --follow ntg-backend-1

database:
	@echo "=== DataBase ==="
	@echo "Building image..."
	@cd database && docker buildx build -t ntg-postgres .
	@echo "Recreating container..."
	@docker compose up -d --no-deps --force-recreate database

database-logs:
	@docker logs --follow ntg-database-1

database-connect:
	@docker exec -it ntg-database-1 psql -h localhost -U changeme -d app

objectstorage:
	@echo "=== ObjectStorage ==="
	@echo "Building image..."
	@cd objectstorage && docker buildx build -t ntg-objectstorage .
	@echo "Recreating container..."
	@docker compose up -d --no-deps --force-recreate objectstorage

objectstorage-logs:
	@docker logs --follow ntg-objectstorage-1

all: backend database objectstorage

up:
	docker compose up -d

down:
	docker compose down

restart: down up
