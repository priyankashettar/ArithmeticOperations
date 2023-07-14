package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	mo "lovoo/src/mathOperations"

	"google.golang.org/grpc"
)

// Initialisation of the variables used for taking command line arguments and the server URI
var (
	serverURI = "localhost:9000"
	method    = flag.String("method", "", "Simple Math Operation you want to work on !")
	a         = flag.Float64("a", 0, "First input variable for the operation")
	b         = flag.Float64("b", 0, "Second input variable for the operation")
)

// Client Dispatcher function to get the result of Math Operation from Server
func mathOperationsDispatcher(client mo.MathOperationsClient, mathOperation string, input mo.Input) {
	switch mathOperation {
	case "add":
		responseMessage, error := client.Add(context.Background(), &input)
		if error == nil {
			fmt.Println(responseMessage.Output)
		} else {
			log.Printf("Server returned an error: %v", error)
		}

	case "sub":
		responseMessage, error := client.Subtract(context.Background(), &input)
		if error == nil {
			fmt.Println(responseMessage.Output)
		} else {
			log.Printf("Server returned an error: %v", error)
		}

	case "mul":
		responseMessage, error := client.Multiply(context.Background(), &input)
		if error == nil {
			fmt.Println(responseMessage.Output)
		} else {
			log.Printf("Server returned an error: %v", error)
		}

	case "div":
		responseMessage, error := client.Division(context.Background(), &input)
		if error == nil {
			fmt.Println(responseMessage.Output)
		} else {
			log.Printf("Server returned an error: %v", error)
		}
	}
}

// Driver function to parse the input arguments and
// subsequently sending the gRPC request to the Server
func main() {

	// Parse the input Arguments
	flag.Parse()

	// Create a connection to gRPC server from client
	clientConnection, error := grpc.Dial(serverURI, grpc.WithInsecure())
	if error != nil {
		panic(error)
	}
	defer clientConnection.Close()
	client := mo.NewMathOperationsClient(clientConnection)

	// Parsed Flag parameters
	mathOperation := *method
	input := mo.Input{X: *a, Y: *b}

	// Call Dispatcher to call math operation from client
	mathOperationsDispatcher(client, mathOperation, input)
}
