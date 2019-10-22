package handler

import (
	"context"

	pb "github.com/steveoc64/petstore/proto"
)

// GetPetByID implements PetstoreService rpc
func (p *PetstoreServer) GetPetByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.Pet, error) {
	return nil, nil
}
