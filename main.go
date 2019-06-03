package main

import (
	"flag"
	"fmt"

	"github.com/infinityworks/email-signature-generator/server"
)

var (
	apiPort int
	version string
)

func init() {
	flag.IntVar(&apiPort, "apiPort", 8080, "Port for the API server")
	flag.StringVar(&version, "version", "0.0.1", "Version of the application currently deployed")
	flag.Parse()
}

func main() {
	a := server.NewApi()
	a.Start(fmt.Sprintf(":%d", apiPort))
}
