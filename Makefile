build:
	@go build -o bin/api
run: build
	@./bin/api
test:
	@go test -v ./...

