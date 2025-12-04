package LLM

func Prompt(user string) ChatRequest {

	prompt := ChatRequest{
		Model: "ministral-3:14b-cloud",
		Messages: []Message{
			{Role: "system", Content: SystemPrompt(Tools())},
			{Role: "user", Content: user},
		},
		Format:  "json",
		Options: Options{Temperature: 0.1},
	}

	return prompt
}

func Ask(prompt string) (*ChatResponse, error) {

	ask := ChatRequest{
		Model: "ministral-3:14b-cloud",
		Messages: []Message{
			{Role: "system", Content: SystemPrompt(Tools())},
			{Role: "user", Content: prompt},
		},
		Format:  "json",
		Options: Options{Temperature: 0.1},
	}

	client := NewOllamaClient()
	resp, err := client.Chat(ask)

	return resp, err
}
