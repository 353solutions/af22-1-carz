package main

import (
	"context"
	"log"
	"net"

	"github.com/353solutions/carz/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterCarzServer(srv, &Server{})
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Update(ctx context.Context, car *pb.Car) (*pb.UpdateResponse, error) {
	log.Printf("Update: %#v\n", car)
	return &pb.UpdateResponse{Id: car.Id}, nil
}

type Server struct {
	pb.UnimplementedCarzServer
}
