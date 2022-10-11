package main

import (
	"context"
	"io"
	"log"

	pb "github.com/fmuyassarov/grpc-go/api/proto"
)

func doGreatManyTimes(c pb.APIServiceClient) {

	log.Println("doGratManyTimes was invoked")
	req := &pb.APIRequest{
		Protocol: "HTTP",
	}

	stream, err := c.APIS(context.Background(), req)

	if err != nil {
		log.Fatal("Failed while calling APIS", err)

	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading the stream", err)
		}

		log.Printf("APIS %s\n", msg)
	}

}
