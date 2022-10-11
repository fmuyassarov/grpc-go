package main

import (
	"fmt"
	"log"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
)

func (s *Server) APIS(in *pb.APIRequest, stream pb.APIService_APISServer) error {
	log.Printf("api stream func was invoked %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s\n, number %d", in.Protocol, i)

		stream.Send(&pb.APIResponse{
			Result: res,
		})

	}
	return nil

}
