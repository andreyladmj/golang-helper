package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	n1 := req.GetFirstNamber()
	n2 := req.GetSecondNumber()
	result := n1 + n2
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {
	fmt.Println("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
