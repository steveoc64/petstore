package handler

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
)

type Database interface {
	GetPetByID(ctx context.Context, id int64) (*pb.Pet, error)
	UpdatePet(ctx context.Context, id int64, name string, status string) error
	AddPet(ctx context.Context, pet *pb.Pet) error
}

// PetstoreServer implements the RPC and REST server
type PetstoreServer struct {
	log      *logrus.Logger
	db       Database
	rpcPort  int
	restPort int
	apiKey   string
}

// NewPetstoreServer returns a new PetstoreServer
func NewPetstoreServer(log *logrus.Logger, db Database, rcpPort, restPort int, apiKey string) *PetstoreServer {
	return &PetstoreServer{
		log:      log,
		db:       db,
		rpcPort:  rcpPort,
		restPort: restPort,
		apiKey:   apiKey,
	}
}

// Run starts and runs the server
func (s *PetstoreServer) Run() {
	s.log.WithField("API_KEY", s.apiKey).Print("Petstore Start Run")
	go s.rpcProxy()
	s.grpcRun()
}
