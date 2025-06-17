package navigation

// Démonstration des fonctionnalités de Navigation Avancée dans les outils JetBrains

import (
	"fmt"
	"math"
	"strings"
)

// ---------------------------------------------------------------------
// 1. Navigation Symbolique - Recherche par symbole (Ctrl+Alt+Shift+N)
// ---------------------------------------------------------------------

// Animal est une interface utilisée pour démontrer la recherche de symboles
// TIP: Placez le curseur sur "Animal" et appuyez sur Alt+F7 pour voir tous ses usages
type Animal interface {
	Nom() string
	FaireBruit() string
}

// Chien implémente l'interface Animal
// TIP: Placez le curseur sur "Animal" dans la ligne ci-dessous et appuyez sur Ctrl+B
// pour naviguer vers la définition de l'interface
type Chien struct {
	nom string
}

func (c Chien) Nom() string {
	return c.nom
}

func (c Chien) FaireBruit() string {
	return "Wouf!"
}

// Chat implémente également l'interface Animal
// TIP: Utilisez Ctrl+Alt+B sur le nom de l'interface "Animal" pour voir toutes ses implémentations
type Chat struct {
	nom string
}

func (c Chat) Nom() string {
	return c.nom
}

func (c Chat) FaireBruit() string {
	return "Miaou!"
}

// ---------------------------------------------------------------------
// 2. Navigation dans la Hiérarchie (Ctrl+H sur un type)
// ---------------------------------------------------------------------

// Forme est une interface permettant de démontrer la navigation hiérarchique
// TIP: Positionnez votre curseur sur "Forme" et appuyez sur Ctrl+H pour voir
// la hiérarchie d'héritage
type Forme interface {
	Aire() float64
	Perimetre() float64
}

// Cercle implémente l'interface Forme
type Cercle struct {
	rayon float64
}

func (c Cercle) Aire() float64 {
	return math.Pi * c.rayon * c.rayon
}

func (c Cercle) Perimetre() float64 {
	return 2 * math.Pi * c.rayon
}

// Rectangle implémente l'interface Forme
type Rectangle struct {
	largeur float64
	hauteur float64
}

func (r Rectangle) Aire() float64 {
	return r.largeur * r.hauteur
}

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur + r.hauteur)
}

// ---------------------------------------------------------------------
// 3. Recherche d'Usages (Alt+F7)
// ---------------------------------------------------------------------

// FormateurTexte est utilisé pour démontrer la recherche d'usages
// TIP: Placez votre curseur sur "FormateurTexte" et appuyez sur Alt+F7
// pour voir tous ses usages dans le projet
type FormateurTexte struct {
	texte string
}

// Mettre en majuscule le texte
func (f *FormateurTexte) EnMajuscules() string {
	return strings.ToUpper(f.texte)
}

// Mettre en minuscules le texte
func (f *FormateurTexte) EnMinuscules() string {
	return strings.ToLower(f.texte)
}

// Première lettre en majuscule
func (f *FormateurTexte) Capitaliser() string {
	if len(f.texte) == 0 {
		return ""
	}
	return strings.ToUpper(f.texte[:1]) + f.texte[1:]
}

// ---------------------------------------------------------------------
// 4. Structure du fichier (Alt+7 ou Ctrl+F12)
// ---------------------------------------------------------------------

// ExempleNavigationStruct est une structure avec plusieurs champs et méthodes
// pour démontrer la vue de structure de fichier
// TIP: Appuyez sur Ctrl+F12 pour afficher la structure de ce fichier
type ExempleNavigationStruct struct {
	Champ1 string
	Champ2 int
	Champ3 bool
	Champ4 []float64
}

// Methode1 est une méthode d'exemple
func (e *ExempleNavigationStruct) Methode1() string {
	return e.Champ1
}

// Methode2 est une méthode d'exemple
func (e *ExempleNavigationStruct) Methode2() int {
	return e.Champ2 * 2
}

// Methode3 est une méthode d'exemple
func (e *ExempleNavigationStruct) Methode3() bool {
	return !e.Champ3
}

// ---------------------------------------------------------------------
// 5. Navigation entre fichiers récemment modifiés (Ctrl+E)
// ---------------------------------------------------------------------

// DemonstrateRecentFiles montre comment naviguer entre fichiers récents
// TIP: Appuyez sur Ctrl+E pour afficher la liste des fichiers récemment ouverts
func DemonstrateRecentFiles() {
	fmt.Println("Appuyez sur Ctrl+E pour voir les fichiers récemment modifiés")
}

// ---------------------------------------------------------------------
// 6. Navigation par numéro de ligne (Ctrl+G)
// ---------------------------------------------------------------------

// DemonstrateGotoLine montre comment aller à une ligne spécifique
// TIP: Appuyez sur Ctrl+G pour aller à une ligne spécifique dans le fichier
func DemonstrateGotoLine() {
	fmt.Println("Appuyez sur Ctrl+G pour aller à une ligne spécifique")
}
