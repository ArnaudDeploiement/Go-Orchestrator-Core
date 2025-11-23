package LLM

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Options  Options   `json:"options"`
	Stream   bool      `json:"stream"`
	Format   Format    `json:"format"`
}

type Format struct {
	Type       any            `json:"type"`
	Properties map[string]any `json:"properties"`
}

type Options struct {
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	MinTokens   int     `json:"min_tokens,omitempty"`
}

type ChatResponse struct {
	Model      string  `json:"model"`
	Message    Message `json:"message"`
	Done       bool    `json:"done"`
	DoneReason string  `json:"done_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
