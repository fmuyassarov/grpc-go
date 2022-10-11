package main

import (
	"log"
	"net"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.APIServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %v\n", addr)
	}
	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterAPIServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start the grpc server %v\n", err)
	}

}
