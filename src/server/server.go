package main

import (
	"context"
	"log"
	"net"

	mo "lovoo/src/mathOperations"

	"google.golang.org/grpc"
)

type Server struct {
	mo.UnimplementedMathOperationsServer
}

// Add receives the protobuff input.
// Performs the Addition operation for the given x, y values and returns the output via gRPC
func (s *Server) Add(ctx context.Context, input *mo.Input) (*mo.Output, error) {
	log.Printf("Adding input variables: %v", input)
	output := Add(input.X, input.Y)
	log.Printf("Returning the sum: %v", output)
	return &mo.Output{
		Output: output,
		Error:  nil,
	}, nil
}

func Add(x, y float64) float64 {
	return x + y
}

// Subtract receives the protobuff input.
// Performs the Subtraction operation for the given x, y values and returns the output via gRPC
func (s *Server) Subtract(ctx context.Context, input *mo.Input) (*mo.Output, error) {
	log.Printf("Subtracting input variables: %v", input)
	output := Subtract(input.X, input.Y)
	log.Printf("Returning the Subtracted value: %v", output)
	return &mo.Output{
		Output: output,
		Error:  nil,
	}, nil
}

func Subtract(x, y float64) float64 {
	return x - y
}

// Multiply receives the protobuff input.
// Performs the Multiplication operation for the given x, y values and returns the output via gRPC
func (s *Server) Multiply(ctx context.Context, input *mo.Input) (*mo.Output, error) {
	log.Printf("Multiplying input variables: %v", input)
	output := Multiply(input.X, input.Y)
	log.Printf("Returning the Multiplied value: %v", output)
	return &mo.Output{
		Output: output,
		Error:  nil,
	}, nil
}

func Multiply(x, y float64) float64 {
	return x * y
}

// Division receives the protobuff input.
// Performs the Division operation for the given x, y values and returns the output via gRPC
func (s *Server) Division(ctx context.Context, input *mo.Input) (*mo.Output, error) {
	log.Printf("Dividing input variables: %v", input)
	output := Division(input.X, input.Y)
	log.Printf("Returning the Divided value: %v", output)
	return &mo.Output{
		Output: output,
		Error:  nil,
	}, nil
}

func Division(x, y float64) float64 {
	return x / y
}

// This is the Driver function of the gPRC Server
func main() {
	log.Println("Starting Math Operations Server!")
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	gRPCServer := grpc.NewServer()
	mo.RegisterMathOperationsServer(gRPCServer, &Server{})

	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
