install-deps:
	go get

.PHONY: test
test:
	go test ./...

build: install-deps test
	go build -ldflags="-s -w" -o build/export cmd/export/main.go; go build -ldflags="-s -w" -o build/workspace cmd/workspace/main.go; cp .togglsheet.example README.md build; zip -r release.zip build