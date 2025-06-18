package coverage_tests

import (
	"errors"
	"strings"
)

// ---------------------------------------------------------------------
// Fonctions pour démontrer les tests et la couverture de code
// ---------------------------------------------------------------------

// CalculerSomme calcule la somme de deux entiers
func CalculerSomme(a, b int) int {
	return a + b
}

// CalculerProduit calcule le produit de deux entiers
func CalculerProduit(a, b int) int {
	return a * b
}

// CalculerQuotient calcule la division de a par b
// Retourne une erreur si b est égal à 0
func CalculerQuotient(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division par zéro")
	}
	return float64(a) / float64(b), nil
}

// VerifierNombrePositif vérifie si un nombre est positif
func VerifierNombrePositif(n int) bool {
	return n > 0
}

// TraiterTexte transforme un texte en majuscules et supprime les espaces
func TraiterTexte(texte string) string {
	texteTraite := strings.ToUpper(texte)
	texteTraite = strings.ReplaceAll(texteTraite, " ", "")
	return texteTraite
}

// ---------------------------------------------------------------------
// Structure pour démontrer les tests de méthodes
// ---------------------------------------------------------------------

// Calculatrice représente une calculatrice simple avec mémoire
type Calculatrice struct {
	memoire int
}

// Ajouter ajoute un nombre à la mémoire
func (c *Calculatrice) Ajouter(n int) {
	c.memoire += n
}

// Soustraire soustrait un nombre de la mémoire
func (c *Calculatrice) Soustraire(n int) {
	c.memoire -= n
}

// ObtenirMemoire retourne la valeur en mémoire
func (c *Calculatrice) ObtenirMemoire() int {
	return c.memoire
}

// ReinitialiserMemoire remet la mémoire à zéro
func (c *Calculatrice) ReinitialiserMemoire() {
	c.memoire = 0
}

// ---------------------------------------------------------------------
// Fonction complexe pour démonstration de couverture partielle
// ---------------------------------------------------------------------

// AnalyserNote évalue une note et retourne un commentaire
func AnalyserNote(note int) string {
	if note < 0 {
		return "Note invalide: nombre négatif"
	}

	switch {
	case note >= 90:
		return "Excellent"
	case note >= 80:
		return "Très bien"
	case note >= 70:
		return "Bien"
	case note >= 60:
		return "Satisfaisant"
	case note >= 50:
		return "Passable"
	default:
		return "Insuffisant"
	}
}

// FiltrerTableau retourne les éléments d'un tableau qui respectent un critère
func FiltrerTableau(elements []int, condition func(int) bool) []int {
	var resultat []int
	for _, element := range elements {
		if condition(element) {
			resultat = append(resultat, element)
		}
	}
	return resultat
}
