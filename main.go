package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example/proto/greet/v1"  // Update this import path to match your project structure
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(), // Make the dial blocking until connection is established or timeout
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new client
	client := pb.NewGreetServiceClient(conn)

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Prepare the request
	request := &pb.GreetRequest{
		Name: "John",
	}

	// Make the gRPC call
	response, err := client.Greet(ctx, request)
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	// Print the response
	fmt.Printf("Response from server: %s\n", response.Greeting)
}
