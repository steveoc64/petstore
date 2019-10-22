package handler

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	pb "github.com/steveoc64/petstore/proto"
)

// GetPetByID implements PetstoreService rpc
func (s *PetstoreServer) GetPetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.Pet, error) {
	s.log.WithFields(logrus.Fields{
		"id": req.Id,
	}).Info("GetPetByID")

	// TODO - get pet using the storace interface

	// lets hardcode one for now
	return &pb.Pet{
		Id: req.Id,
		Category: &pb.Category{
			Id:   1,
			Name: "cat",
		},
		Name:      fmt.Printf("Pet number %v", req.Id),
		PhotoUrls: nil,
		Tags: []*pb.Tag{
			&pb.Tag{
				Id:   1,
				Name: "Easy to deal with",
			},
			&pb.Tag{
				Id:   2,
				Name: "House Trained",
			},
		},
		Status: pb.Status_available.String(),
	}, nil
}
