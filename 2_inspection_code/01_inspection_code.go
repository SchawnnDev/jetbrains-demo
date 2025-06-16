package inspection

// Démonstration des Inspections de Code dans GoLand
// GoLand offre une analyse statique puissante qui détecte les problèmes potentiels
// avant même l'exécution du code.

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 1. Exemple d'inspection de variables non utilisées
func VariablesNonUtilisees() {
	// GoLand surlignera cette variable car elle n'est jamais utilisée
	compteur := 0

	// Cette variable est utilisée
	message := "Hello"
	fmt.Println(message)
}

// 2. Exemple d'inspection de code mort (unreachable code)
func CodeInatteignable() bool {
	// GoLand détectera que le code après le return ne sera jamais exécuté
	return true
	fmt.Println("Ce code ne sera jamais exécuté")
	return false
}

// 3. Exemple d'inspection des ressources non fermées
func RessourcesNonFermees() {
	// GoLand détectera que le fichier n'est jamais fermé
	fichier, err := os.Open("exemple.txt")
	if err != nil {
		return
	}

	donnees := make([]byte, 100)
	_, err = fichier.Read(donnees)
	// Manque defer fichier.Close() - GoLand le signalera
}

// 4. Exemple d'inspection de gestion d'erreur incomplète
func GestionErreurIncomplete() string {
	// GoLand détectera que l'erreur n'est pas vérifiée
	contenu, _ := os.ReadFile("config.json")
	return string(contenu)
}

// 5. Exemple d'inspection de variables qui pourraient être constantes
func VariablesConstantes() {
	// GoLand suggérera de transformer cette variable en constante
	piApprox := 3.14
	fmt.Println("Pi est approximativement", piApprox)
}

// 6. Exemple d'inspection de simplifications possibles
func SimplificationsCode(condition bool) int {
	// GoLand suggérera de simplifier cette structure conditionnelle
	if condition {
		return 1
	} else {
		return 0
	}
}

// 7. Exemple d'inspection de problèmes de performance
func ProblemePerformance(texte string) []string {
	resultat := []string{}

	// GoLand suggérera d'allouer une capacité initiale au slice
	for _, mot := range strings.Split(texte, " ") {
		resultat = append(resultat, mot)
	}

	return resultat
}

// 8. Exemple d'inspection de fonctions non testées
func FonctionNonTestee(a, b int) int {
	// GoLand peut identifier les fonctions sans tests unitaires correspondants
	return a + b
}

// 9. Exemple d'inspection des copies coûteuses
func CopiesCouteuses(data []byte) {
	// GoLand détectera qu'une grande structure est copiée plutôt que passée par référence
	reader := strings.NewReader(string(data))
	CopieReader(reader)
}

func CopieReader(r io.Reader) {
	// La copie de r pourrait être coûteuse si r est grande
	// GoLand suggérera de passer par pointeur
}

// 10. Exemple d'inspection de conversions redondantes
func ConversionsRedondantes() {
	nombre := 42
	// GoLand détectera cette conversion redondante
	autreNombre := int(nombre)
	fmt.Println(autreNombre)
}
