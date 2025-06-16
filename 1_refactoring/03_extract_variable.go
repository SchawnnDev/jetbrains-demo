package refactoring

import (
	"fmt"
	"math"
)

// Démonstration du Refactoring d'Extraction de Variable
// Vous pouvez extraire des expressions en variables pour une meilleure lisibilité
// Pour démontrer:
// 1. Sélectionnez l'expression que vous souhaitez extraire
// 2. Appuyez sur Ctrl+Alt+V (Windows/Linux) ou Commande+Alt+V (macOS)
// 3. Entrez un nom pour la nouvelle variable et appliquez

// SurfaceCercle calcule la surface d'un cercle sans variables extraites
func SurfaceCercle(rayon float64) {
	// Ici, l'expression math.Pi * rayon * rayon pourrait être extraite
	fmt.Printf("La surface du cercle de rayon %.2f est %.2f\n",
		rayon, math.Pi*rayon*rayon)

	// Si nous avons besoin de la surface à nouveau, nous devons répéter le calcul
	if math.Pi*rayon*rayon > 100 {
		fmt.Println("C'est un grand cercle")
	}
}

// ExpressionComplexe possède une expression complexe qui pourrait être extraite
func ExpressionComplexe(a, b, c int) int {
	// Cette expression complexe pourrait être extraite pour rendre le code plus lisible
	return (a*a + b*b) * (c * c) / (a + b + c)
}
