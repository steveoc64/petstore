package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
	"google.golang.org/grpc"
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

// grpcRun runs the RPC server
// Looks ugly, but its just common boilerplate that would normally be in a lib
func (s *PetstoreServer) grpcRun() {
	endpoint := fmt.Sprintf(":%d", s.rpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		s.log.WithError(err).Fatal("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(grpcServer, s)

	s.log.WithField("endpoint", endpoint).Println("Serving gRPC")

	grpcServer.Serve(lis)
}

// rpcProxy hooks up the REST endpoints.
// Looks ugly, but its just common boilerplate that would normally be in a lib
func (s *PetstoreServer) rpcProxy() error {
	// Use our custom error handler
	runtime.HTTPError = CustomHTTPError
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	rpcendpoint := fmt.Sprintf(":%d", s.rpcPort)
	webendpoint := fmt.Sprintf(":%d", s.restPort)
	err := pb.RegisterPetstoreServiceHandlerFromEndpoint(ctx, mux, rpcendpoint, opts)
	if err != nil {
		return err
	}

	s.log.WithField("endpoint", webendpoint).Println("Serving REST Proxy")
	return http.ListenAndServe(webendpoint, mux)
}

type errorBody struct {
	Err string `json:"error,omitempty"`
}

// CustomHTTPError for stripping error contents back
func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	contentType := marshaler.ContentType()
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Del("Trailer")

	errMsg := s.Message()
	statusCode := runtime.HTTPStatusFromCode(status.Code(err))

	// Examine leader on the message, and use that for the custom error code
	if len(errMsg) > 3 && errMsg[3] == ':' {
		v, err := strconv.Atoi(errMsg[:3])
		if err == nil {
			statusCode = v
			errMsg = errMsg[4:]
		}
	}
	e := errorBody{Err: errMsg}
	w.WriteHeader(statusCode)
	jErr := json.NewEncoder(w).Encode(e)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}
