build:
	@go build -o bin/api
run: build
	@./bin/api
lint:
	@golangci-lint run
test:
	@go test -v ./...

