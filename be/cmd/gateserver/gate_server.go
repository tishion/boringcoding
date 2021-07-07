package main

import (
	"flag"
	"fmt"

	"mmc.com/landingtask/be/cmd/gateserver/servercore"
	"mmc.com/landingtask/be/internal/common"
)

// declares the arguments flags
var (
	port = flag.Int(
		"p",
		8081,
		"The port to use")

	rpcAddress = flag.String(
		"r",
		"localhost:5100",
		"The RPC service connection string")

	cacheAddress = flag.String(
		"c",
		"localhost:6379",
		"The cache service connetion string")
)

func main() {
	// parse the arguments
	flag.Parse()

	// create and start the server
	server := servercore.CreateServer(*rpcAddress, *cacheAddress)
	endpoint := fmt.Sprintf("0.0.0.0:%d", *port)
	common.LogI("Starting gate_server on: %s", endpoint)
	server.Start(endpoint)
}
