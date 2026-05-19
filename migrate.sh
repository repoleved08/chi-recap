#!/bin/bash

# Migration helper script for chi-recap

set -e

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | xargs)
fi

DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-3306}
DB_USERNAME=${DB_USERNAME:-root}
DB_PASSWORD=${DB_PASSWORD:-rootpass}
DB_NAME=${DB_NAME:-chi_recap}

case "$1" in
    "up")
        echo "Running migrations up..."
        go run main.go migrate -up
        ;;
    "down")
        echo "Running migrations down..."
        go run main.go migrate -down
        ;;
    "status")
        echo "Checking migration status..."
        go run main.go migrate -status
        ;;
    "create")
        if [ -z "$2" ]; then
            echo "Usage: ./migrate.sh create <migration_name>"
            exit 1
        fi
        TIMESTAMP=$(date +%s)
        # Create migration files with correct timestamp format
        MIGRATION_NUM=$(ls migrations/ 2>/dev/null | grep -o '^[0-9]*' | sort -n | tail -1 || echo 0)
        NEXT_NUM=$((MIGRATION_NUM + 1))
        
        UP_FILE="migrations/$(printf '%06d' $NEXT_NUM)_${2}.up.sql"
        DOWN_FILE="migrations/$(printf '%06d' $NEXT_NUM)_${2}.down.sql"
        
        echo "-- Migration: $2" > "$UP_FILE"
        echo "-- Write your migration SQL here" >> "$UP_FILE"
        echo "" >> "$UP_FILE"
        
        echo "-- Rollback: $2" > "$DOWN_FILE"
        echo "-- Write your rollback SQL here" >> "$DOWN_FILE"
        echo "" >> "$DOWN_FILE"
        
        echo "✓ Created migration files:"
        echo "  - $UP_FILE"
        echo "  - $DOWN_FILE"
        ;;
    "fresh")
        echo "Dropping all tables and running fresh migrations..."
        go run main.go migrate -down
        go run main.go migrate -up
        echo "✓ Database refreshed"
        ;;
    "help"|"")
        echo "Migration management script"
        echo ""
        echo "Usage: ./migrate.sh [command]"
        echo ""
        echo "Commands:"
        echo "  up              - Run all pending migrations"
        echo "  down            - Revert the last migration"
        echo "  status          - Show current migration status"
        echo "  create <name>   - Create a new migration file"
        echo "  fresh           - Revert and re-run all migrations"
        echo "  help            - Show this help message"
        ;;
    *)
        echo "Unknown command: $1"
        echo "Run './migrate.sh help' for usage information"
        exit 1
        ;;
esac
