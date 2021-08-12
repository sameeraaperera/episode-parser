run:
	go run cmd/main.go

test: ## Runs the tests
	go test -v ./...
