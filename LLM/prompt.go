package LLM

func Prompt(user string) ChatRequest {

	prompt := ChatRequest{
		Model: "llama3.1:8b",
		Messages: []Message{
			{Role: "system", Content: SystemPrompt(Tools())},
			{Role: "user", Content: user},
		},
		Format:  "json",
		Options: Options{Temperature: 0.1},
	}

	return prompt
}
