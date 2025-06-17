package inspection

// Démonstration des Inspections de Code dans GoLand
// GoLand offre une analyse statique puissante qui détecte les problèmes potentiels
// avant même l'exécution du code.

import (
	"fmt"
	"strings"
)

// 1. Exemple d'inspection de variables non utilisées
func VariablesNonUtilisees() {
	// GoLand surlignera cette variable car elle n'est jamais utilisée
	compteur := 0

	// Cette variable est utilisée
	message := "Hello"
	fmt.Println(message)
}

// 2. Exemple d'inspection de code mort (unreachable code)
func CodeInatteignable() bool {
	// GoLand détectera que le code après le return ne sera jamais exécuté
	return true
	fmt.Println("Ce code ne sera jamais exécuté")
	return false
}

// 3. Exemple d'inspection de simplifications possibles
func SimplificationsCode(condition bool) int {
	// GoLand suggérera de simplifier cette structure conditionnelle
	if condition {
		return 1
	} else {
		return 0
	}
}
