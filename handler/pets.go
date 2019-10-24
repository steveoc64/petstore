package handler

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"

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

// UpdatePetWithForm updates the name and status of a Pet using form encoded data, converted to protobuf
func (s *PetstoreServer) UpdatePetWithForm(ctx context.Context, req *pb.UpdatePetWithFormReq) (*pb.Empty, error) {
	s.log.WithFields(logrus.Fields{
		"id":     req.PetId,
		"name":   req.Name,
		"status": req.Status,
	}).Info("UpdatePet")

	err := s.db.UpdatePetWithForm(ctx, req.PetId, req.Name, req.Status)
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
	// get the passed in APIKey and validate first
	var apiKey string
	if headers, ok := metadata.FromIncomingContext(ctx); ok {
		apiKeys := headers["api_key"]
		if len(apiKeys) < 1 || apiKeys[0] != s.apiKey {
			return nil, fmt.Errorf("400:Invalid API_KEY Supplied")
		}
		apiKey = apiKeys[0]
	}
	s.log.WithFields(logrus.Fields{
		"id":      req.PetId,
		"api_key": apiKey,
	}).Info("DeletePet")
	err := s.db.DeletePet(ctx, req.PetId)
	return &pb.Empty{}, err
}

// FindPetsByStatus gets all the pets that match any of the passed in statuses
func (s *PetstoreServer) FindPetsByStatus(ctx context.Context, req *pb.StatusReq) (*pb.Pets, error) {
	s.log.WithField("statuses", strings.Join(req.Status, ",")).Info("FindPetsByStatus")
	return &pb.Pets{}, nil
}

// UpdatePet updates a pet from the input data
func (s *PetstoreServer) UpdatePet(ctx context.Context, req *pb.Pet) (*pb.Pet, error) {
	s.log.WithField("id", req.PetId).Info("UpdatePet")
	err := s.db.UpdatePet(ctx, req)
	if err != nil {
		return nil, err
	}
	// In the swaggerAPI example, calling this REST endpoint returns the updated pet details
	// we do the same here
	return s.db.GetPetByID(ctx, req.PetId)
}

// UploadFile uploads a photo against a pet
func (s *PetstoreServer) UploadFile(ctx context.Context, req *pb.UploadFileReq) (*pb.ApiResponse, error) {
	s.log.WithField("id", req.PetId).Info("UploadFile")
	err := s.db.UploadFile(ctx, req.PetId, req.File)
	if err != nil {
		return &pb.ApiResponse{
			Code:    1,
			Type:    "error",
			Message: fmt.Sprintf("upload error %s", err.Error()),
		}, err
	}
	return &pb.ApiResponse{
		Code:    11,
		Type:    "type 11",
		Message: "api success response of type 11",
	}, nil
}
