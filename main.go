package main

import (
	"github.com/steveoc64/petstore/handler"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	defaultRpcPort  = 8081
	defaultRestPort = 8080
)

func main() {
	var (
		rpcPort, restPort int
		apiKey            string
		err               error
	)

	// Setup logrus
	log := logrus.New()
	log.SetReportCaller(true)
	log.Print("Starting Petstore")

	// Get the runtime params from the ENV vars
	rpcPort, err = strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		rpcPort = defaultRpcPort
		log.Error("Missing RPC_PORT")
	}

	restPort, err = strconv.Atoi(os.Getenv("REST_PORT"))
	if err != nil {
		restPort = defaultRestPort
		log.Error("Missing REST_PORT")
	}
	apiKey = os.Getenv("API_KEY"))

	petstore := handler.NewPetstoreServer(log, rpcPort, restPort)
}
