.PHONY: setup-dev build-protoc build-client run-server test-unit 

# Install grpc
requirements:
	go get -u google.golang.org/grpc

# Initial setup requirements that installs protoc plugins
setup-dev: requirements
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Builds the project from .proto file
build-protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	src/mathOperations/mathOperations.proto

# Builds the client
build-client:
	go build -o client src/client/client.go

# Start the Math Operations Server
run-server:
	go run src/server/server.go

# Run the unit tests
test-unit:
	go test src/server/* -v

