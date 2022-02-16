package main

import (
	"context"
	"log"
	"os"

	// "os"

	pb "github.com/restlesswhy/grps-passwordgener/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// sample  = "exezz"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPasswordGeneratorServiceClient(conn)

	sample := os.Args[1]

	resp, err := client.Generate(context.Background(), &pb.PasswordRequest{
		Sample: sample,
	})
	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	log.Println("New password:", string(resp.GetPassword()))
}
