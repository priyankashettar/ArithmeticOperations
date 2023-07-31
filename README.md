# Basic Arithmetic Operater
This is a gRPC command line application which enables the user to do the basic Math Operations.
The application has a gRPC Server and a client.
The gRPC Server supports the basic Math Operations of two variables. It presently supports
    - Addition
    - Subtraction
    - Multiplication and 
    - Division operations

## Overview
- [Requirements](#requirements)
- [Running locally](#running-locally)
- [Developing the Application](#developing-the-application)
- [References](#references)


### Requirements
Make sure to have Golang and Protobuf in your system.
- Install the latest version of (golang)[https://go.dev/doc/install]
- Install the (protobuf)[https://grpc.io/docs/protoc-installation/]
- Install the `grpc` package. Make file can help you here `make requirements`

### Running Locally
Once the above requirements are met the user can use this application.
Below are the steps involved while using the application
- Build the client
- Start the server in a terminal and 
- Make a call from the client to this server from another terminal


#### Build the Client
The user needs to build the client application. There is a make command to help you here !
``` bash
make build-client
```

#### Start the Math Operator Server
Use your terminal to start the server application. Use the below command
``` bash
make run-server
```

#### Use Client
Presently the server supports addition, subtraction, multiplication and division of two variables.
Below are the example commands which can help you in using the client for getting results for above operations.
``` bash
./client --help
./client add -a 8 -b 2
./client sub -a 8 -b 2
./client mul -a 8 -b 2
./client div -a 8 -b 2
```
- `Addition` operation is done by `add` method
- `Subtraction` operation is done by `sub` method
- `Multiplication` operation is done by `mul` method
- `Division` operation is done by `div` method
- `a` and `b` are the two input variables to be provided

### Developing the Application
Below steps can be used for enhancing this application. The steps involved are,
- Make sure to have system  requirements
- Install the protoc plugins
- Once you make changes in code make sure to write the test cases and test them
- Build the code and run the client/server for enhancements

#### Install protoc plugins
``` bash
make build-protoc
```
- Make sure to export the plugins in path `export PATH="$PATH:$(go env GOPATH)/bin"`

#### Run the Unit Tests
``` bash
make test-unit
```

### References
- https://grpc.io/docs/languages/go/quickstart/
- https://go.dev/doc/
