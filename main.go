package main

import (
	"fmt"
	"orchestrator/server"
)

func main() {
	fmt.Println("Starting Orchestrator...")
	srv := server.NewServer(":8080")
	srv.Start()

}
