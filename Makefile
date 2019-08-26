.PHONY: default
default: standalone 

.PHONY: standalone
standalone:
	go run cmd/standalone/main.go

.PHONY: templating
templating:
	go run cmd/templating/main.go


