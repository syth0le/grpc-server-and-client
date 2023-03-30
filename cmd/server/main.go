package main

import (
	"fmt"
	"github.com/syth0le/grpc-server-and-client/pkg/pb"
	"github.com/syth0le/grpc-server-and-client/pkg/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	pb.RegisterEchoServer(s, srv)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v; %v", s.GetServiceInfo(), listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
