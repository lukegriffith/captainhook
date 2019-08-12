.PHONY: default
default: prereq test build


.PHONY: prereq
prereq:
	go get . 

.PHONY: build
build: 
	go build -o ./build/hook .

.PHONY: test
test: 
	go test ./...




