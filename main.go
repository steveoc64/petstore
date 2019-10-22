package main

import "github.com/sirupsen/logrus"

func main() {

	log := logrus.New()
	log.SetReportCaller(true)
	log.Print("Starting Petstore")

}
