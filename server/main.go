package main

import (
	"log"
	"net"

	"github.com/restlesswhy/grps-passwordgener/proto"
	"github.com/restlesswhy/grps-passwordgener/server/passwordservice"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	service := new(passwordservice.PasswordServer)
	grpcsrv := grpc.NewServer()
	proto.RegisterPasswordGeneratorServiceServer(grpcsrv, service)
	grpcsrv.Serve(listener)
}