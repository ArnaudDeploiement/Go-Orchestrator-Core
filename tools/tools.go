package tools

import "fmt"

func Motivate(param string) string {
	return fmt.Sprintf("Je voudrais activer l'outil motivation suite à ton propos: %s", "'"+param+"'"+"."+"\n Mais Arnaud doit l'implémenter. Au travail !")
}

func GenerateAnswer(param string) string {
	return fmt.Sprintf("Answer generation tool activated %s ", param)
}

func InverseClass(param string) string {
	return fmt.Sprintf("Je voudrais activer l'outil inversion de classe suite à ton propos : %s, mais Arnaud doit l'implémenter.", param)
}

func Blague(param string) string {
	return fmt.Sprintf("Je voudrais activer l'outil blague suite à ton propos : %s, mais Arnaud doit l'implémenter.", param)
}
