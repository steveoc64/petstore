package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	pb "github.com/steveoc64/petstore/proto"
)

// GetPetByID implements PetstoreService rpc
func (s *PetstoreServer) GetPetByID(ctx context.Context, req *pb.PetID) (*pb.Pet, error) {
	// TODO - implement metrics (instead of just logging)
	// TODO - add handlers to manage context timeouts
	s.log.WithFields(logrus.Fields{
		"id": req.PetId,
	}).Info("GetPetByID")

	pet, err := s.db.GetPetByID(ctx, req.PetId)
	if err != nil {
		return nil, err
	}
	if pet == nil {
		return nil, fmt.Errorf("404:Pet %#v not found", req.PetId)
	}
	return pet, nil
}

// UpdatePet updates the name and status of a Pet
func (s *PetstoreServer) UpdatePetWithForm(ctx context.Context, req *pb.UpdatePetWithFormReq) (*pb.Empty, error) {
	s.log.WithFields(logrus.Fields{
		"id":     req.PetId,
		"name":   req.Name,
		"status": req.Status,
	}).Info("UpdatePet")

	err := s.db.UpdatePet(ctx, req.PetId, req.Name, req.Status)
	return &pb.Empty{}, err
}

// AddPet adds a pet to the store
func (s *PetstoreServer) AddPet(ctx context.Context, req *pb.Pet) (*pb.Pet, error) {
	s.log.WithFields(logrus.Fields{
		"id":   req.PetId,
		"name": req.Name,
	}).Info("AddPet")

	// We are accepting plain JSON here - so unmarshal it into a Pet struct
	//pet := &pb.Pet{}
	//if err := json.Unmarshal([]byte(req.Body), pet); err != nil {
	//s.log.WithField("body", req.Body).WithError(err).Error("Unmarshal body error")
	//return nil, errors.New("405:Invalid Input")
	//}
	pet := req

	err := s.db.AddPet(ctx, pet)
	if err != nil {
		return nil, err
	}
	return s.db.GetPetByID(ctx, req.PetId)
}

// DeletePet removes a pet. Check the req header for the API_KEY value
func (s *PetstoreServer) DeletePet(ctx context.Context, req *pb.DeletePetReq) (*pb.Empty, error) {
	s.log.WithField("id", req.PetId).Info("DeletePet")
	return &pb.Empty{}, nil
}

// FindPetsByStatus gets all the pets that match any of the passed in statuses
func (s *PetstoreServer) FindPetsByStatus(ctx context.Context, req *pb.StatusReq) (*pb.Pets, error) {
	s.log.WithField("statuses", strings.Join(req.Status, ",")).Info("FindPetsByStatus")
	return &pb.Pets{}, nil
}

// UpdatePet updates a pet from the input data
func (s *PetstoreServer) UpdatePet(ctx context.Context, req *pb.Pet) (*pb.Empty, error) {
	s.log.WithField("id", req.PetId).Info("UpdatePet")
	return &pb.Empty{}, nil
}

// UploadFile uploads a photo against a pet
func (s *PetstoreServer) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.Empty, error) {
	s.log.WithField("id", req.PetId).Info("UploadFile")
	return &pb.Empty{}, nil
}
