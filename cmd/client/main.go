package main

import (
	"context"
	"flag"
	"github.com/syth0le/grpc-server-and-client/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("Not enough arguments")
	}
	name := flag.Arg(0)

	age, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal("Not enough arguments")
	}

	ctx := context.Background()
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewEchoClient(conn)
	res, err := client.SendMessage(ctx, &pb.MessageRequest{
		Name: name,
		Age:  int32(age),
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetResult())
}
