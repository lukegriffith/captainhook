.PHONY: default
default: test build


.PHONY: build
build: 
	go build -o ./build/hook .

.PHONY: test
test: 
	go test ./...




