package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fmuyassarov/grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50050"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("failed to connect to the server %v\n", err)
	}

	c := pb.NewCalculatorServiceClient(conn)

	// goSum(c)
	goPrime(c)

	defer conn.Close()
}

// func goSum(c pb.CalculatorServiceClient) {
// 	log.Print("goSum was invoked")
// 	res, err := c.API(context.Background(), &pb.CalculatorRequest{
// 		X: 10,
// 		Y: 25,
// 	})
// 	if err != nil {
// 		log.Fatal("failed to request sum of 10 and 25 from the server", err)
// 	}
// 	log.Printf("Received response: %v\n", res.Sum)

// }

func goPrime(c pb.CalculatorServiceClient) {
	log.Print("goPrime was invoked")
	stream, err := c.Primes(context.Background(), &pb.PrimeRequest{
		PrimeNumber: 400,
	})
	if err != nil {
		log.Fatal("failed to request prime number from server", err)

	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error while reading the stream", err)
		}
		log.Printf("received stream value from server is %v: ", msg.Sum)
	}

}
