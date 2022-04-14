package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/abrahamSN/shippy/shippy-service-vessel/proto/vessel"
	"go-micro.dev/v4"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// find available vessel
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

// Our grpc service handler
type vesselService struct {
	repo Repository
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{
			Id:        "vessel001",
			Capacity:  500,
			MaxWeight: 200000,
			Name:      "Boaty McBoatface",
			Available: false,
			OwnerId:   "",
		},
	}
	repo := &VesselRepository{vessels}

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), &vesselService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
