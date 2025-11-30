package LLM

func Prompt(user string) ChatRequest {

	prompt := ChatRequest{
		Model: "mistral-nemo:latest",
		Messages: []Message{
			{Role: "system", Content: SystemPrompt(Tools())},
			{Role: "user", Content: user},
		},
		Format:  "json",
		Options: Options{Temperature: 0.1},
	}

	return prompt
}
