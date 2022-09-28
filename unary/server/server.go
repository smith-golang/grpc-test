package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/smith-golang/grpc-test/unary/unarypb"
	"google.golang.org/grpc"
)

type server struct {
	unarypb.UnimplementedGreetingServiceServer
}

func (*server) Greet(ctx context.Context, req *unarypb.GreetRequest) (*unarypb.GreetResponse, error) {
	fmt.Println("Greet functions was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &unarypb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("------Unary Testing--------")

	lis, err := net.Listen("tcp", "0.0.0.0:50051") //default port for gRPC
	if err != nil {
		log.Fatal("Failed to listen :%v", err)
	}

	s := grpc.NewServer()
	unarypb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to served %v", err)
	}
}