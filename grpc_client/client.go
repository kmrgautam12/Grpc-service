package main

import (
	"context"
	pb "grpc/protos"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	grpcClient, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Error occured creating grpc client")
	}
	cc := pb.NewGetRecordClient(grpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := cc.PersonRecord(ctx, &pb.PersonRecordRequest{
		Name:  "Gautam",
		Id:    123,
		Email: "gautam.kumar",
	},
	)
	if err != nil {
		log.Fatal("Error making grpc call:", err)
	}
	log.Println("grpc response is", resp)
}
