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


.PHONY: dbtest
dbtest:
	go -o ./build/dbtest database.go
