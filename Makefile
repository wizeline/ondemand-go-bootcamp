APP_BIN = capstone

# Test controller layer
test-controller:
	go test ./internal/controller

# Test service layer
test-service:
	go test ./internal/service

# Test repository layer
test-repo:
	go test ./internal/repository

# Test the clean architecture's layers: controller, service and repository.
test-clean-architecture:test-controller test-service test-repo

# Run all tests
test-all:
	go test ./...

# Build the application
build:
	go mod download
	go mod verify
	go mod tidy -v
	go build -o $(APP_BIN) cmd/main.go

run: build
	./$(APP_BIN)

clean:
	@[ -f '$(APP_BIN)' ] && rm -v $(APP_BIN) || true

