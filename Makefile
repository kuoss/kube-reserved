.PHONY: test test
test:
	go test -v ./...

.PHONY: test lint
lint:
	golangci-lint run
