package agent

import "log"

var Incoming = make(chan string, 1) // buffer de 1 pour éviter le deadlock simple

func RespEndpoint(content string) string {
	// Tu crées ta payload pour le tool
	payload := content

	// Envoi NON bloquant (ou dans une goroutine)
	select {
	case Incoming <- payload:
		// ok
	default:
		// channel plein, log ou gérer le cas
		log.Println("agent.Incoming is full, dropping payload")
	}

	return payload
}
