APP_NAME=web
DB_USERNAME=root
DB_PASSWORD=root1234
DB_NAME=scholar_ai
DB_PORT=3306

.PHONY: dev build test migrate up down

# Run in development mode
dev:
	export GIN_MODE=debug && go run ./cmd/$(APP_NAME)

# Build the application
build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

# Run tests
test:
	go test -v ./...


# Atlas Migration Commands
# All commands auto-generate timestamped migration names (format: YYYYMMDDHHMMSS)
# Apply all pending migrations
migrate:
	atlas migrate diff --env gorm
up:
	atlas schema apply --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"
down:
	atlas migrate down --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"
clean:
	atlas schema clean --env gorm -u "mysql://${DB_USERNAME}:${DB_PASSWORD}@localhost:3306/${DB_NAME}"