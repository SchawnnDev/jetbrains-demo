package refactoring

// Démonstration du Refactoring de Déplacement
// Vous pouvez déplacer des fonctions, méthodes ou champs entre packages et types
// Pour démontrer:
// 1. Sélectionnez la fonction/méthode/champ que vous souhaitez déplacer
// 2. Appuyez sur F6 ou cliquez-droit et sélectionnez Refactoriser > Déplacer
// 3. Sélectionnez la destination et appliquez

// UtilisateurAvantDeplacement démontre une struct avec des méthodes avant déplacement
type UtilisateurAvantDeplacement struct {
	ID     int
	Prenom string
	Nom    string
	Email  string
}

// ValiderEmail méthode qui pourrait être déplacée vers un package utilitaire
func (u *UtilisateurAvantDeplacement) ValiderEmail() bool {
	// Validation simple pour la démonstration
	return len(u.Email) > 0 && (u.Email[0] != '@' && u.Email[len(u.Email)-1] != '@')
}

// NomComplet méthode qui pourrait être déplacée vers un autre type
func (u *UtilisateurAvantDeplacement) NomComplet() string {
	return u.Prenom + " " + u.Nom
}
