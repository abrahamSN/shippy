package main

import (
	"context"
	"log"
	"os"

	pb "github.com/abrahamSN/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/abrahamSN/shippy/shippy-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags.
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	h := &handler{repository, vesselClient}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterShippingServiceHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
