package refactoring

// Démonstration du Refactoring de Renommage
// Vous pouvez renommer des variables, fonctions, structs, interfaces et packages
// Pour démontrer:
// 1. Placez le curseur sur un identifiant
// 2. Appuyez sur Maj+F6 ou cliquez-droit et sélectionnez Refactoriser > Renommer
// 3. Entrez le nouveau nom et appliquez

// RenommeMoi représente une struct qui pourrait être renommée
type RenommeMoi struct {
	AncienChamp string
	Compteur    int
}

// AncienNomDeFonction peut être renommé
func AncienNomDeFonction(ancienParametre string) string {
	ancienneVariable := "Cette variable pourrait être renommée"
	return ancienneVariable + " " + ancienParametre
}

// Utilisation démontre l'utilisation des composants qui pourraient être renommés
func Utilisation() {
	instance := RenommeMoi{AncienChamp: "valeur", Compteur: 5}
	resultat := AncienNomDeFonction(instance.AncienChamp)
	_ = resultat
}
