BIN=bin
REDIS_SERVER=redis_server

default: test build
run: test build execute

build:
	CGO_ENABLED=0 go build -o $(BIN)/$(REDIS_SERVER) ./cmd/server/

test:
	# Run go test with cache disabled
	CGO_ENABLED=0 go test -cover ./... -count=1

execute:
	$(BIN)/$(REDIS_SERVER) -addr=127.0.0.1:1234