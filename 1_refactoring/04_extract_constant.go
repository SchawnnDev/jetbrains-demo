package refactoring

import (
	"fmt"
	"time"
)

// Démonstration du Refactoring d'Extraction de Constante
// Vous pouvez extraire des valeurs littérales en constantes nommées
// Pour démontrer:
// 1. Sélectionnez la valeur littérale que vous souhaitez extraire
// 2. Appuyez sur Ctrl+Alt+C (Windows/Linux) ou Commande+Alt+C (macOS)
// 3. Entrez un nom pour la nouvelle constante et appliquez

// Configuration montre du code avec des nombres magiques et des chaînes
func Configuration() {
	// Ces nombres magiques et ces chaînes pourraient être extraits en constantes
	timeout := 30 * time.Second
	maxRetries := 5
	endpoint := "https://api.example.com/v1/data"

	fmt.Printf("Connexion à %s avec un timeout de %v et %d essais\n",
		endpoint, timeout, maxRetries)

	// Si nous avons besoin de ces valeurs à nouveau, nous devons les dupliquer
	if connexionEchouee() {
		fmt.Printf("Tentative de reconnexion à %s jusqu'à %d fois\n",
			"https://api.example.com/v1/data", 5)
	}
}

// Fonction auxiliaire pour l'exemple
func connexionEchouee() bool {
	return true // Juste pour la démonstration
}
