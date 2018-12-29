install-deps:
	go get

.PHONY: test
test:
	go test ./...

build: install-deps test
	go build -o build/export cmd/export/main.go; go build -o build/workspace cmd/workspace/main.go; cp example.config.toml README.md build