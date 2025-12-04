package agent

func Idle(a *Agent) StateFN {

	a.State = "je ne fais rien"

	return Idle
}

func Evaluate(a *Agent) StateFN {

	a.State = "je suis en train d'évaluer"

	return Evaluate
}
