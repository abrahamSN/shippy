package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/abrahamSN/shippy/shippy-service-consignment/proto/consignment"
	"go-micro.dev/v4"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	// setup connection to the server
	service := micro.NewService(micro.Name("shippy.cli.consignment"))
	service.Init()

	client := pb.NewShippingService("shippy.service.consignment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, errCreate := client.CreateConsignment(context.Background(), consignment)
	if errCreate != nil {
		log.Fatalf("Could not greet: %v", errCreate)
	}
	log.Printf("Created: %t", r.Created)

	getAll, errGet := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if errGet != nil {
		log.Fatalf("Could not list consignments: %v", errGet)
	}

	for i, v := range getAll.Consignments {
		log.Println(i, v)
	}
}
