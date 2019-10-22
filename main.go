package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()
	//log.SetFormatter(&log.JSONFormatter{})
	log.Print("Starting Petstore")

}
