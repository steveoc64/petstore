package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/steveoc64/petstore/database/memory"
	"github.com/steveoc64/petstore/database/mysql"
	"github.com/steveoc64/petstore/database/testdb"
	"github.com/steveoc64/petstore/handler"
)

const (
	defaultRPCPort  = 8081
	defaultRestPort = 8080
)

func main() {
	var (
		rpcPort, restPort int
		apiKey            string
		err               error
	)

	// create the photos directory if it doesnt already exist
	os.Mkdir("photos", 0777)

	// Setup logrus
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.Print("Starting Petstore")

	// Get the runtime params from the ENV vars
	rpcPort, err = strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		rpcPort = defaultRPCPort
		log.Error("Missing RPC_PORT")
	}

	restPort, err = strconv.Atoi(os.Getenv("REST_PORT"))
	if err != nil {
		restPort = defaultRestPort
		log.Error("Missing REST_PORT")
	}
	apiKey = os.Getenv("API_KEY")

	var db handler.Database
	dbname := strings.ToUpper(os.Getenv("DATABASE"))
	switch dbname {
	case "MEMORY":
		db = memory.New()
	case "TESTDB":
		db = testdb.New()
		log.Info("Created TESTDB with pets [1..5]")
	case "MYSQL":
		dsn := os.Getenv("DSN")
		db, err = mysql.New(log, dsn)
		if err != nil {
			log.WithError(err).WithField("dsn", dsn).Fatal("Error opening mysql connection")
		}
	default:
		log.Errorf("Invalid DATABASE value %#v", dbname)
		db = memory.New()
	}
	log.WithField("database", dbname).Info("Connected to DB")

	petstore := handler.NewPetstoreServer(log, db, rpcPort, restPort, apiKey)
	petstore.Run()
}
