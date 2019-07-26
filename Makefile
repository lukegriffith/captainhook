.PHONY: default
default: test build


.PHONY: build
build: 
	go build -o ./build/hooks .

.PHONY: test
test: 
	go test ./...




