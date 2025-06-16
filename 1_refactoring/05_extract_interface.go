package refactoring

// Démonstration du Refactoring d'Extraction d'Interface
// Vous pouvez créer automatiquement une interface à partir d'un type existant
// Pour démontrer:
// 1. Placez le curseur sur un nom de type
// 2. Appuyez sur Ctrl+Maj+Alt+T (Windows/Linux) ou Ctrl+T (macOS) pour ouvrir le menu de refactoring
// 3. Sélectionnez "Extraire l'Interface"
// 4. Sélectionnez les méthodes à inclure et nommez l'interface

// StockageFichiers est un type concret dont nous pourrions vouloir extraire une interface
type StockageFichiers struct {
	cheminBase string
}

// Sauvegarder enregistre des données dans un fichier
func (sf *StockageFichiers) Sauvegarder(cle string, donnees []byte) error {
	// Détails d'implémentation...
	return nil
}

// Charger récupère des données depuis un fichier
func (sf *StockageFichiers) Charger(cle string) ([]byte, error) {
	// Détails d'implémentation...
	return nil, nil
}

// Supprimer supprime un fichier
func (sf *StockageFichiers) Supprimer(cle string) error {
	// Détails d'implémentation...
	return nil
}

// ObtenirCheminBase renvoie le chemin de base (non nécessaire dans l'interface)
func (sf *StockageFichiers) ObtenirCheminBase() string {
	return sf.cheminBase
}
