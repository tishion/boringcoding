package main

import (

	//local packages
	"flag"
	"fmt"

	"mmc.com/landingtask/be/cmd/userserver/servercore"
	"mmc.com/landingtask/be/internal/common"
)

// declares the arguments flags
var (
	port = flag.Int(
		"p",
		5100,
		"The port to use")

	db = flag.String(
		"d",
		"test:test@tcp(localhost:3306)/test",
		"The database service connection string")
)

func main() {
	// parse the arguments
	flag.Parse()

	// create and start the server
	server := servercore.CreateServer(*db)
	endpoint := fmt.Sprintf("0.0.0.0:%d", *port)
	common.LogI("Starting user_server on: %s", endpoint)
	server.Start(endpoint)
}
