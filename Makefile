APP_BIN = capstone

test:
	go test ./...

build:
	go build -o $(APP_BIN) cmd/main.go

run: build
	./$(APP_BIN)

clean:
	@[ -f '$(APP_BIN)' ] && rm -v $(APP_BIN) || true