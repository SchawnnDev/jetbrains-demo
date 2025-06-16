package refactoring

import "fmt"

// Démonstration du Refactoring d'Inlining
// Vous pouvez inliner des variables, constantes ou fonctions
// Pour démontrer:
// 1. Placez le curseur sur la variable, constante ou appel de fonction que vous souhaitez inliner
// 2. Appuyez sur Ctrl+Alt+N (Windows/Linux) ou Commande+Alt+N (macOS)
// 3. Confirmez l'opération d'inlining

// Exemple avec des variables:

// AvantInliningVariable démontre du code avant l'inlining de variable
func AvantInliningVariable(valeur int) {
	// La variable résultat pourrait être inlinée
	resultat := valeur * 2
	fmt.Println("Le résultat est:", resultat)
}

// Exemple avec des fonctions:

// FonctionAuxiliaire qui pourrait être inlinée
func FonctionAuxiliaire(x int) int {
	return x * x
}

// AvantInliningFonction démontre du code avant l'inlining de fonction
func AvantInliningFonction(valeur int) {
	// L'appel de fonction pourrait être inliné
	resultat := FonctionAuxiliaire(valeur)
	fmt.Println("Carré:", resultat)
}
