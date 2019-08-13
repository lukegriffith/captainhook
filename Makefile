.PHONY: default
default: 
	echo "run make standalone"

.PHONY: standalone
standalone:
	go get ./build/standalone
	go build -o ./build/standalone cmd/standalone/


