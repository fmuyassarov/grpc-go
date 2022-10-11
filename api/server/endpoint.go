package main

import (
	"context"
	"log"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
)

// RPC endpoint
func (s *Server) API(ctx context.Context, in *pb.APIRequest) (*pb.APIResponse, error) {
	log.Printf("API function was invoked with %v\n", in)
	return &pb.APIResponse{
		Result: "Hello" + in.Protocol,
	}, nil
}
