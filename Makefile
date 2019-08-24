.PHONY: default
default: standalone 

.PHONY: standalone
standalone:
	go run cmd/standalone/main.go


