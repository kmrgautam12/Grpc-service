package main

import (
	"context"
	p "grpc/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	p.UnimplementedGetRecordServer
}

func (s *server) PersonRecord(ctx context.Context, req *p.PersonRecordRequest) (*p.PersonRecordResponse, error) {
	return &p.PersonRecordResponse{
		Name:  req.Name,
		Id:    "12",
		Email: req.Email,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	p.RegisterGetRecordServer(grpcServer, &server{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Error in serving on grpc port", err)
	}
	log.Println("Grpc serving on port:", "50051")
}
