package server

import (
	"context"
	"fmt"
	"github.com/syth0le/grpc-server-and-client/pkg/pb"
)

type GRPCServer struct{}

func (s *GRPCServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{
		Result: fmt.Sprintf("Name: %s, age: %v. ITS RESULT", req.GetName(), req.GetAge()),
	}, nil
}
