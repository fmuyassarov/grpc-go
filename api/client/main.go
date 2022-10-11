package main

import (
	"log"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to the server %v\n", err)
	}

	c := pb.NewAPIServiceClient(conn)

	// doGreet(c)
	doGreatManyTimes(c)

	defer conn.Close()

}
