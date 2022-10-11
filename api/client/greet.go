package main

import (
	"context"
	"log"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
)

func doGreet(c pb.APIServiceClient) {
	log.Print("doGreet was invoked")
	res, err := c.API(context.Background(), &pb.APIRequest{
		Protocol: "TCP",
	})

	if err != nil {
		log.Fatalf("could not send the greetings %v\n", err)
	}

	log.Printf("Received response: %s\n", res.Result)

}
