
.PHONY: test

test:
	go generate
	goimports -w .
	go fmt ./...
	go test ./...