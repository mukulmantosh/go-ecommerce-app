build:
	@go build -o bin/api
run: build
	@./bin/api
lint:
	@golangci-lint run
clean:
	@go clean -modcache
test:
	@go test -v ./... -count=1
cover:
	@go test -v ./... -covermode=atomic -coverpkg=./... -coverprofile coverage.out
	@go tool cover -func=coverage.out -o=coverage.out
