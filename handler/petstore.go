package handler

import (
	"context"

	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
)

// Database defines the set of methods that the petstore server needs to be able to treat any object as a database
type Database interface {
	GetPetByID(ctx context.Context, id int64) (*pb.Pet, error)
	UpdatePetWithForm(ctx context.Context, id int64, name string, status string) error
	DeletePet(ctx context.Context, id int64) error
	UploadFile(ctx context.Context, id int64, url string) error
	AddPet(ctx context.Context, pet *pb.Pet) error
	UpdatePet(ctx context.Context, pet *pb.Pet) error
	FindPetsByStatus(ctx context.Context, statuses []string) (*pb.Pets, error)
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
