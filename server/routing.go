package server

import (
	"encoding/json"
	"net/http"
	"orchestrator/LLM"
)

func router(content []byte, tc *LLM.Tool, apiResp *LLM.APIResponse, w http.ResponseWriter) {

	if err := json.Unmarshal(content, &tc); err == nil && tc.Tool != "" {

		apiResp.Tool = tc.Tool

		switch apiResp.Tool {
		case "motivate_tool":
			apiResp.Result = LLM.Motivate(tc.Params)

		case "generate_answer_tool":
			apiResp.Result = LLM.GenerateAnswer(tc.Params)

		case "inverse_class_tool":
			apiResp.Result = LLM.InverseClass(tc.Params)

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
