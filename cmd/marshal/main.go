package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/353solutions/carz/pb"
)

func main() {
	loc := pb.Location{
		Lat: -22.97192,
		Lng: -43.21050,
	}
	fmt.Println(loc.String())

	data, err := proto.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("proto:", len(data))

	jdata, err := json.Marshal(loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("json:", len(jdata))

	car := pb.Car{
		Id:       "007",
		DriverId: "Bond",
		Location: &loc,
		Time:     timestamppb.Now(),
		Status:   pb.Status_FREE,
	}
	fmt.Println(car.String())
	data, err = proto.Marshal(&car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("car proto:", len(data))
	jdata, err = json.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("car json:", len(jdata))
	fmt.Println(string(jdata))

	jdata, err = protojson.Marshal(&car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("car protojson:", len(jdata))
	fmt.Println(string(jdata))
}
