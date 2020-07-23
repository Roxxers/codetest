package main

import (
	"flag"
	"fmt"

	"thirdlight.com/aggregation-server/server"
)

var defaultPort uint = 8000

func main() {
	server := server.SetupRouter()
	port := flag.Uint("p", defaultPort, "Port to run aggregation server on")
	server.Run(fmt.Sprintf(":%d", *port))
}
