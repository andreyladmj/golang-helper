package main

import (
	pb "../api"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address backend")
	output := flag.String("o", "output.wav", "wav file where the output will be written")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("too low arguments")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend)
	if err != nil {
		log.Fatal("could not connect to the backend")
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn, grpc.WithInsecure())

	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("could not say %s: %v", text.Text, err)
	}
	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("could not writer to file %s: %v", *output, err)
	}
}
