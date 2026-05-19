.PHONY: help build run db-up db-down db-logs test clean

help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make run        - Run the application"
	@echo "  make db-up      - Start MySQL database"
	@echo "  make db-down    - Stop MySQL database"
	@echo "  make db-logs    - View database logs"
	@echo "  make db-shell   - Connect to MySQL shell"
	@echo "  make deps       - Download dependencies"
	@echo "  make clean      - Remove binary"
	@echo "  make test       - Run tests"

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
	docker exec -it chi-recap-mysql mysql -u root -proot chi_recap

deps:
	go mod download
	go mod tidy

clean:
	rm -f chi-recap

test:
	go test -v ./...
