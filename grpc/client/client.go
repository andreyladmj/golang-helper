package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("hello client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to the server: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{FirstNamber: 7, SecondNumber: 9}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.SumResult)
}
