.PHONY: help build run db-up db-down db-logs test clean migrate-status migrate-seed

help:
	@echo "Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  make build           - Build the application"
	@echo "  make run             - Run the application (starts DB automatically)"
	@echo "  make deps            - Download dependencies"
	@echo "  make clean           - Remove binary"
	@echo ""
	@echo "Database:"
	@echo "  make db-up           - Start MySQL database"
	@echo "  make db-down         - Stop MySQL database"
	@echo "  make db-logs         - View database logs"
	@echo "  make db-shell        - Connect to MySQL shell"
	@echo ""
	@echo "Migrations:"
	@echo "  make migrate-status  - Show migration status"
	@echo "  make migrate-seed    - Seed database with sample data"
	@echo ""
	@echo "Testing:"
	@echo "  make test            - Run tests"

build:
	go build -o chi-recap

run: db-up
	go run main.go

db-up:
	docker-compose up -d
	@echo "Waiting for database to be ready..."
	@sleep 5

db-down:
	docker-compose down

db-logs:
	docker-compose logs -f mysql

db-shell:
	docker exec -it chi-recap-mysql mysql -u root -prootpass chi_recap

deps:
	go mod download
	go mod tidy

clean:
	rm -f chi-recap

test:
	go test -v ./...

migrate-status:
	go run main.go migrate -status

migrate-seed:
	go run main.go migrate -seed
