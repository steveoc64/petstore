package handler

import (
	"context"
	"fmt"
	"net"
	"net/http"

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

func (s *PetstoreServer) grpcRun() {
	endpoint := fmt.Sprintf(":%d", s.rpcPort)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		s.log.WithError(err).Fatal("failed to listen")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPetstoreServiceServer(grpcServer, s)

	s.log.WithFields(logrus.Fields{
		"port":     s.rpcPort,
		"endpoint": endpoint,
	}).Println("Serving gRPC")

	grpcServer.Serve(lis)
}

func (s *PetstoreServer) rpcProxy() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	rpcendpoint := fmt.Sprintf(":%d", s.rpcPort)
	webendpoint := fmt.Sprintf(":%d", s.restPort)
	err := pb.RegisterPetstoreServiceServer(ctx, mux, rpcendpoint, opts)
	if err != nil {
		return err
	}

	s.log.WithFields(logrus.Fields{
		"port":     s.restPort,
		"endpoint": webendpoint,
	}).Println("Serving REST Proxy")
	return http.ListenAndServe(webendpoint, mux)
}
