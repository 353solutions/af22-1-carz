package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/353solutions/carz/pb"
)

func main() {
	car := &pb.Car{
		Id:       "007",
		DriverId: "Bond",
		Location: &pb.Location{
			Lat: 51.4871871,
			Lng: -0.1266743,
		},
		Time:   timestamppb.Now(),
		Status: pb.Status_FREE,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	addr := "localhost:8888"
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to %s", addr)

	c := pb.NewCarzClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	resp, err := c.Update(ctx, car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s updated\n", resp.Id)
}
