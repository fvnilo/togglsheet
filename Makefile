install-deps:
	go get

.PHONY: test
test:
	go test ./...

build: install-deps test
	go build -ldflags="-s -w" -o build/togglsheet cmd/togglsheet/main.go; cp .togglsheet.example README.md build; zip -r release.zip build