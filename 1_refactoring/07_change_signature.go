package refactoring

import "fmt"

// Démonstration du Refactoring de Changement de Signature
// Vous pouvez modifier les signatures de fonctions (paramètres et valeurs de retour)
// Pour démontrer:
// 1. Placez le curseur sur le nom de la fonction
// 2. Appuyez sur Ctrl+F6 (Windows/Linux) ou Commande+F6 (macOS)
// 3. Utilisez la boîte de dialogue pour ajouter/supprimer/réorganiser les paramètres et changer les valeurs de retour

// AvantChangementSignature montre une fonction avant de changer sa signature
func AvantChangementSignature(nom string) string {
	return fmt.Sprintf("Bonjour, %s !", nom)
}

// Utilisation de la fonction originale
func UtilisationAvantChangement() {
	message := AvantChangementSignature("Monde")
	fmt.Println(message)
}
