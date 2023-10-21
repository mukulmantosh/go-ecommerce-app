build:
	@go build -o bin/api
run: build
	@./bin/api
lint:
	@golangci-lint run
test:
	@go test -v ./...
cover:
	@go test -v ./... -covermode=atomic -coverpkg=./... -coverprofile coverage.out
	@go tool cover -func=coverage.out -o=coverage.out
