.PHONY: default
default: run


.PHONY: test
test: 
	go build . 
	go test ./...


*.go:
	go test .
	go run main.go

run:
	go run main.go
