package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"

	pb "github.com/steveoc64/petstore/proto"
)

// GetPetByID implements PetstoreService rpc
func (s *PetstoreServer) GetPetByID(ctx context.Context, req *pb.ID) (*pb.Pet, error) {
	// TODO - implement metrics (instead of just logging)
	// TODO - add handlers to manage context timeouts
	s.log.WithFields(logrus.Fields{
		"id": req.Id,
	}).Info("GetPetByID")

	pet, err := s.db.GetPetByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if pet == nil {
		return nil, fmt.Errorf("404:Pet %#v not found", req.Id)
	}
	return pet, nil
}

// UpdatePet updates the name and status of a Pet
func (s *PetstoreServer) UpdatePet(ctx context.Context, req *pb.UpdatePetRequest) (*pb.Empty, error) {
	s.log.WithFields(logrus.Fields{
		"id":     req.Id,
		"name":   req.Name,
		"status": req.Status,
	}).Info("Update pet")

	err := s.db.UpdatePet(ctx, req.Id, req.Name, req.Status)
	return &pb.Empty{}, err
}

// AddPet adds a pet to the store
func (s *PetstoreServer) AddPet(ctx context.Context, req *pb.Pet) (*pb.Pet, error) {
	s.log.WithFields(logrus.Fields{
		"id":   req.Id,
		"name": req.Name,
	}).Info("Adding pet")
	if req.Id == 0 {
		grpc.SendHeader(ctx, metadata.MD{
			"foo":        []string{"foo"},
			"statuscode": []string{"405"},
		})
		return nil, fmt.Errorf("405:Invalid ID 0")
	}

	err := s.db.AddPet(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.db.GetPetByID(ctx, req.Id)
}
