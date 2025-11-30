package agent

func NewAgent(name string, systemPrompt string, initialBehavior StateFN, registryTool RegistryTool) *Agent {
	return &Agent{
		Name:         name,
		SystemPrompt: systemPrompt,
		StateFN:      initialBehavior,
		RegistryTool: registryTool,
	}
}

func (a *Agent) Run() {
	for a.StateFN != nil {
		a.StateFN = a.StateFN(a)
	}
}
