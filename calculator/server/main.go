package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fmuyassarov/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50050"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %v\n", addr)
	}
	log.Printf("Listening on %v\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to start the grpc server %v\n", err)
	}

}

func (s *Server) API(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Sum: in.X + in.Y,
	}, nil
}
