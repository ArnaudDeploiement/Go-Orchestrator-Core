package server

import (
	"encoding/json"
	"log"
	"net/http"
	"orchestrator/LLM"
	"orchestrator/agent"
	"orchestrator/tools"
	"time"
)

func router(apiResp *LLM.APIResponse, w http.ResponseWriter) {
	var val string
	select {
	case val = <-agent.Incoming:
		// ok
	case <-time.After(5 * time.Second):
		apiResp.Tool = "fallback"
		apiResp.Result = "timeout waiting for agent"
		return
	}

	var tool tools.Tool
	if err := json.Unmarshal([]byte(val), &tool); err != nil {
		log.Printf("failed to unmarshal tool payload %q: %v", val, err)
		apiResp.Tool = "fallback"
		apiResp.Result = "invalid tool payload"
		return
	}

	switch tool.Tool {
	case "motivate_tool":
		apiResp.Result = tools.Motivate(tool.Params)

	case "generate_answer_tool":
		apiResp.Result = tools.GenerateAnswer(tool.Params)

	case "inverse_class_tool":
		apiResp.Result = tools.InverseClass(tool.Params)

	case "no_tool":
		apiResp.Result = tools.NoTool(tool.Params)

	default:
		http.Error(w, "unknown tool: "+tool.Tool, http.StatusBadRequest)
		return
	}
}
