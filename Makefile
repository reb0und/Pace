BINARY_NAME=Pace

all: build run

# Based on https://gist.github.com/rcmachado/af3db315e31383502660
.PHONY: help
## Display this help text
help:
	$(info Available Targets)
	@awk '/^[a-zA-Z\-\_0-9]+:/ {                    \
		nb = sub( /^## /, "", helpMsg );              \
		if(nb == 0) {                                 \
		helpMsg = $$0;                              \
		nb = sub( /^[^:]*:.* ## /, "", helpMsg );   \
		}                                             \
		if (nb)                                       \
		printf "\033[1;31m%-" width "s\033[0m %s\n", $$1, helpMsg;   \
	}                                               \
	{ helpMsg = $$0 }'                              \
	$(MAKEFILE_LIST) | column -ts:

.PHONY: build
## Build the binary
build:
	mkdir -p bin
	go build -o bin/${BINARY_NAME} cmd/main.go

.PHONY: run
## Run the binary
run:
	./bin/pace

clean:
	go clean
	rm ./bin/${BINARY_NAME}

vet:
	go vet ./...

dep:
	go mod download

lint:
	gofumpt -w cmd & gofumpt -w internal & gofumpt -w configs & golangci-lint run --disable-all -E errcheck