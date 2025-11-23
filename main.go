package main

import (
	"orchestrator/server"
)

func main() {

	srv := server.NewServer(":8080")
	srv.Start()

}
