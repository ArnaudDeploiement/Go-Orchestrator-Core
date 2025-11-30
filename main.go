package main

import (
	"fmt"
	"orchestrator/agent"
	"orchestrator/server"
	"orchestrator/tools"
)

func main() {
	fmt.Println("Starting Orchestrator...")
	srv := server.NewServer(":8080")
	srv.Start()

	agent := agent.NewAgent("Agent 1", "Tu es agent qui fait des blagues", agent.Idle, agent.RegistryTool{Tools: []tools.Tool{{Tool: "blague"}}})

	go agent.Run()
	agent.Message = "Hello, world!"

}
