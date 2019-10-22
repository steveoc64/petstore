package handler

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	pb "github.com/steveoc64/petstore/proto"
	"google.golang.org/grpc"
)

type PetstoreServer struct {
	log      *logrus.Logger
	rpcPort  int
	restPort int
	apiKey   string
}

func NewPetstoreServer(log *logrus.Logger, rcpPort, restPort int, apiKey string) *PetstoreServer {
	return &PetstoreServer{
		log:      log,
		rpcPort:  rcpPort,
		restPort: restPort,
		apiKey:   apiKey,
	}
}

func (s *PetstoreServer) grpcRun(log *logrus.Logger) {

	endpoint := fmt.Sprintf(":%d", s.rpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.WithError(err).Fatal("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(grpcServer, &petStoreServer{})

	log.WithFields(logrus.Fields{
		"port":     port,
		"endpoint": endpoint,
	}).Println("Serving gRPC")

	grpcServer.Serve(lis)
}

func rpcProxy(log *logrus.Logger) error {
	var (
		port int
		err  error
	)

	port, err = strconv.Atoi(os.Getenv("REST_PORT"))
	if err != nil {
		port = defaultRestPort
		log.Error("Missing REST_PORT")
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	rpcendpoint := fmt.Sprintf(":%d", s.port)
	webendpoint := fmt.Sprintf(":%d", s.web)
	err := rp.RegisterGameServiceHandlerFromEndpoint(ctx, mux, rpcendpoint, opts)
	if err != nil {
		return err
	}

	s.log.WithFields(logrus.Fields{
		"port":     s.web,
		"endpoint": webendpoint,
	}).Println("Serving REST Proxy")
	return http.ListenAndServe(webendpoint, mux)
}
