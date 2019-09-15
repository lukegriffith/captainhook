.PHONY: default
default: standalone 

.PHONY: standalone
standalone:
	GODEBUG=http2debug=2 go run cmd/standalone/main.go

.PHONY: templating
templating:
	go run cmd/templating/main.go

.PHONY: test
test:
	go test .

