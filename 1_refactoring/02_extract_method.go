package refactoring

import "fmt"

// Démonstration du Refactoring d'Extraction de Méthode/Fonction
// Vous pouvez extraire des blocs de code en fonctions séparées
// Pour démontrer:
// 1. Sélectionnez le bloc de code que vous souhaitez extraire
// 2. Appuyez sur Ctrl+Alt+M (Windows/Linux) ou Commande+Alt+M (macOS)
// 3. Entrez un nom pour la nouvelle méthode/fonction et appliquez

// TraiterDonnees montre du code avant l'extraction de méthode
func TraiterDonnees(donnees []int) {
	// Cette fonction contient du code qui pourrait être extrait en fonctions séparées

	// Calculer la somme - cela pourrait être extrait
	somme := 0
	for _, valeur := range donnees {
		somme += valeur
	}
	fmt.Println("Somme:", somme)

	// Calculer la moyenne - cela pourrait être extrait
	moyenne := 0
	if len(donnees) > 0 {
		moyenne = somme / len(donnees)
	}
	fmt.Println("Moyenne:", moyenne)

	// Trouver le maximum - cela pourrait être extrait
	max := 0
	if len(donnees) > 0 {
		max = donnees[0]
		for _, valeur := range donnees {
			if valeur > max {
				max = valeur
			}
		}
	}
	fmt.Println("Maximum:", max)
}
