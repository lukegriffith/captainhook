.PHONY: default
default: test build


.PHONY: build
build: 
	go build .

.PHONY: test
test: 
	go test ./...




