package agent

import "orchestrator/tools"

type Agent struct {
	Name         string       `json:"name"`
	SystemPrompt string       `json:"system_prompt"`
	Memory       Memory       `json:"memory"`
	State        string       `json:"state"`
	StateFN      StateFN      `json:"-"`
	Goal         Goal         `json:"goal"`
	Plan         Plan         `json:"plan"`
	Message      string       `json:"message"`
	RegistryTool RegistryTool `json:"registry_tool"`
}

type Memory struct {
	Experience []string `json:"experience"`
}

type Plan struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	Action string `json:"action"`
	Data   []Data `json:"data"`
}

type Data struct {
	Context string `json:"context"`
}

type StateFN func(*Agent) StateFN

type Goal struct {
	Description string `json:"description"`
}

type RegistryTool struct {
	Tools []tools.Tool `json:"tools"`
}
