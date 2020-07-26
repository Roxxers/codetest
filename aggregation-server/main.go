package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	"thirdlight.com/aggregation-server/server"
)

var defaultPort uint = 8000

func main() {
	server := server.SetupRouter()
	port := flag.Uint("p", defaultPort, "Port to run aggregation server on")
	debug := flag.Bool("d", false, "Run in debug mode")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	server.Run(fmt.Sprintf(":%d", *port))
}
