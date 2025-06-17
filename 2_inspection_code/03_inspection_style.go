package inspection

// Démonstration des Inspections de Style et Formatage dans GoLand
// GoLand propose des inspections qui aident à maintenir un code cohérent et lisible.

import (
	"fmt"
)

// 1. Exemple d'inspection de formatage non conforme
func FormatageNonConforme() {
	// GoLand détectera l'indentation incorrecte et les espaces manquants
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 2. Exemple d'inspection de nommage non conforme
func nommage_non_conforme() {
	// GoLand signalera que les noms de fonctions en Go suivent généralement le style camelCase
	// et non snake_case
	VAR_GLOBALE := "valeur"
	fmt.Println(VAR_GLOBALE)
}

// 3. Exemple d'inspection de strings qui pourraient utiliser des raw strings
func StringsNonOptimisees() string {
	// GoLand suggérera d'utiliser une raw string pour éviter les échappements multiples
	chemin := "C:\\Users\\nom\\Documents\\fichiers\\données.txt"
	return chemin
}

// 4. Exemple d'inspection de déclarations redondantes
func DeclarationsRedondantes() {
	// GoLand signalera la redondance des types qui peuvent être inférés
	var nombre int = 42 // Le ": int" est redondant avec le littéral 42
	fmt.Println(nombre)
}

// . Exemple d'inspection de blocs imbriqués excessifs
func BlocsImbriquesExcessifs(a, b, c int) {
	// GoLand peut suggérer de simplifier des structures de contrôle trop imbriquées
	if a > 0 {
		if b > 0 {
			if c > 0 {
				fmt.Println("Tous positifs")
			} else {
				fmt.Println("A et B positifs")
			}
		} else {
			fmt.Println("Seulement A positif")
		}
	} else {
		fmt.Println("A négatif")
	}
}
