package server

import (
	"encoding/json"
	"net/http"
	"orchestrator/LLM"
	"orchestrator/tools"
)

func router(content []byte, tc *tools.Tool, apiResp *LLM.APIResponse, w http.ResponseWriter) {

	if err := json.Unmarshal(content, &tc); err == nil && tc.Tool != "" {

		apiResp.Tool = tc.Tool

		switch apiResp.Tool {
		case "motivate_tool":
			apiResp.Result = tools.Motivate(tc.Params)

		case "generate_answer_tool":
			apiResp.Result = tools.GenerateAnswer(tc.Params)

		case "inverse_class_tool":
			apiResp.Result = tools.InverseClass(tc.Params)

		case "no_tool":
			apiResp.Result = tc.Params

		default:
			http.Error(w, "unknown tool: "+tc.Tool, http.StatusBadRequest)
			return
		}
	} else {
		apiResp.Tool = "fallback"
		apiResp.Result = "Error" + tc.Params
	}

}
