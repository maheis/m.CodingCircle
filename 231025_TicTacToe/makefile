default: build

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: analyze test
	@mkdir -p ./build
	@go build -o ./build/tictactoe

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

test: analyze
	@go test -cover ./...

.PHONY: analyze build coverage test 