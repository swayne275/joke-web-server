package main

import (
	"fmt"

	"github.com/swayne275/joke-web-server/api"
)

const (
	// port used to host the client API
	port = "5000"
)

func main() {
	fmt.Printf("Starting joke api server on port %s\n", port)
	err := api.StartServer(port)
	if err != nil {
		panic(err)
	}
}
