package inspection

// Démonstration des Inspections de Style et Formatage dans GoLand
// GoLand propose des inspections qui aident à maintenir un code cohérent et lisible.

import (
	"fmt"
	"strings"
)

// 1. Exemple d'inspection de formatage non conforme
func FormatageNonConforme(){
  // GoLand détectera l'indentation incorrecte et les espaces manquants
  for i:=0;i<10;i++{
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

// 3. Exemple d'inspection de commentaires manquants pour les fonctions exportées
// GoLand signalera que cette fonction exportée (commençant par une majuscule)
// devrait avoir un commentaire
func FonctionExportee(valeur int) int {
	return valeur * 2
}

// 4. Exemple d'inspection des imports non utilisés
import (
	"time"  // GoLand signalera que ce package est importé mais non utilisé
)

// 5. Exemple d'inspection d'imports non organisés
import "os"
// GoLand suggérera de grouper les imports

// 6. Exemple d'inspection de strings qui pourraient utiliser des raw strings
func StringsNonOptimisees() string {
	// GoLand suggérera d'utiliser une raw string pour éviter les échappements multiples
	chemin := "C:\\Users\\nom\\Documents\\fichiers\\données.txt"
	return chemin
}

// 7. Exemple d'inspection de déclarations redondantes
func DeclarationsRedondantes() {
	// GoLand signalera la redondance des types qui peuvent être inférés
	var nombre int = 42  // Le ": int" est redondant avec le littéral 42
	fmt.Println(nombre)
}

// 8. Exemple d'inspection de constantes magiques
func ConstantesMagiques(taille int) bool {
	// GoLand suggérera d'extraire cette valeur magique en constante nommée
	return taille > 1024
}

// 9. Exemple d'inspection de blocs imbriqués excessifs
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

// 10. Exemple d'inspection de longueur excessive de fonctions
func FonctionTropLongue() {
	// GoLand peut avertir si une fonction est trop longue et devrait être découpée
	// Imaginez ici une fonction de plusieurs centaines de lignes...
	fmt.Println("Cette fonction devrait être découpée en plusieurs fonctions plus petites")
	// ... beaucoup de code ...

	// Les lignes commentées simulent une fonction très longue
	// pour les besoins de la démonstration
	// ...
	// ...
	// ...
}
