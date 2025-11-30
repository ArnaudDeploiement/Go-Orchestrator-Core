package agent

func Idle(a *Agent) StateFN {

	a.State = "je ne fais rien"

	return Idle
}
