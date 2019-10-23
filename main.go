package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/steveoc64/petstore/database/memory"
	"github.com/steveoc64/petstore/database/mysql"
	"github.com/steveoc64/petstore/database/testdb"
	"github.com/steveoc64/petstore/handler"
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
	log.SetFormatter(&logrus.JSONFormatter{})
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
	apiKey = os.Getenv("API_KEY")

	var db handler.Database
	dbname := os.Getenv("DATABASE")
	switch dbname {
	case "MEMORY":
		db = memory.NewMemoryDB()
	case "TESTDB":
		db = testdb.NewTestDB()
	case "MYSQL":
		db, err = mysql.NewMysqlDB(log, os.Getenv("DSN"))
	default:
		log.Errorf("Invalid DATABASE value %#v", dbname)
		db = memory.NewMemoryDB()
	}
	log.WithField("database", dbname).Info("Connected to DB")

	petstore := handler.NewPetstoreServer(log, db, rpcPort, restPort, apiKey)
	petstore.Run()
}
