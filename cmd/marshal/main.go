package main

import (
	"encoding/json"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

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
}
