package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Démonstration des fonctionnalités de Debugging et Profiling dans GoLand
// GoLand offre des outils puissants pour déboguer et profiler les applications Go

// ---------------------------------------------------------------------
// PARTIE 1: DÉBOGUEUR
// ---------------------------------------------------------------------

// 1. Points d'arrêt et inspection des variables
func CalculationAvecBug(a, b int) int {
	// TIP: Placez un point d'arrêt ici en cliquant sur le côté gauche de cette ligne
	// ou en appuyant sur <Ctrl+F8>

	// Cette fonction contient un bug à trouver avec le débogueur
	result := 0
	for i := 0; i < b; i++ {
		// TIP: Pendant le débogage, survolez les variables pour voir leurs valeurs
		if i%2 == 0 {
			result += a
		} else {
			// Bug: cette ligne soustrait a au lieu de l'ajouter quand i est impair
			result -= a
		}
	}

	return result
}

// 2. Évaluation d'expressions pendant le débogage
func CalculerSommeTableau(tableau []int) int {
	// TIP: Pendant le débogage, vous pouvez évaluer des expressions personnalisées
	// en utilisant "Evaluate Expression" (Alt+F8)

	somme := 0
	for i, valeur := range tableau {
		_ = i
		// TIP: Survolez 'somme' pendant le débogage pour voir sa valeur changer
		// TIP: Placez un point d'arrêt conditionnel ici (par exemple: i > 5)
		somme += valeur
	}

	return somme
}

// 3. Points d'arrêt conditionnels et watchpoints
func ProcessArray(array []int) []int {
	result := make([]int, 0, len(array))

	for i, val := range array {
		// TIP: Ajoutez un point d'arrêt conditionnel ici (clic droit sur le point
		// d'arrêt > Condition > val > 100)
		if val%3 == 0 {
			// TIP: Ajoutez un watchpoint sur result (clic droit > Add to Watches)
			result = append(result, val*2)
		} else if val%5 == 0 {
			result = append(result, val*3)
		} else {
			result = append(result, val)
		}

		// TIP: Utilisez Step Over (F8) pour avancer d'une ligne
		fmt.Printf("Traitement de l'élément %d à l'index %d\n", val, i)
	}

	return result
}

// 4. Déboguer une goroutine
func DeboguageGoroutines() {
	// TIP: GoLand permet de déboguer facilement les goroutines
	// Dans la fenêtre de débogage, vous pouvez voir toutes les goroutines

	resultChan := make(chan int)

	// Lancement de plusieurs goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			// TIP: Placez un point d'arrêt ici pour voir le débogueur s'arrêter
			// dans une goroutine
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			resultChan <- id * 10
		}(i)
	}

	// Récupération des résultats
	for i := 0; i < 5; i++ {
		result := <-resultChan
		fmt.Printf("Reçu: %d\n", result)
	}
}

// ---------------------------------------------------------------------
// PARTIE 2: PROFILER
// ---------------------------------------------------------------------

// 1. Fonction avec problème de performance
func FonctionLente() {
	// TIP: Cette fonction a des problèmes de performance à identifier avec le profiler
	// Utilisez CPU Profiler (Run > Profile 'nom_fonction' with CPU Profiler)

	tableau := genererGrandTableau(100000)

	// Version inefficace du tri
	trierTableauInefficace(tableau)

	// Calcul inefficace de la somme
	calculerSommeInefficace(tableau)
}

// Génère un grand tableau aléatoire
func genererGrandTableau(taille int) []int {
	tableau := make([]int, taille)
	for i := 0; i < taille; i++ {
		tableau[i] = rand.Intn(10000)
	}
	return tableau
}

// Version inefficace du tri - à optimiser
func trierTableauInefficace(tableau []int) {
	// TIP: Le profiler montrera que cette implémentation est inefficace
	// Comparer avec la version optimisée

	n := len(tableau)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if tableau[j] > tableau[j+1] {
				tableau[j], tableau[j+1] = tableau[j+1], tableau[j]
			}
		}
	}
}

// Version optimisée du tri
func trierTableauOptimise(tableau []int) {
	// Utilisation du package sort standard
	sort.Ints(tableau)
}

// Calcul inefficace de la somme - à optimiser
func calculerSommeInefficace(tableau []int) int {
	// TIP: Cette fonction est inefficace car elle calcule la somme de manière répétitive
	somme := 0
	for i := 0; i < len(tableau); i++ {
		// Calcul répétitif inutile
		for j := 0; j <= i; j++ {
			if j == i {
				somme += tableau[j]
			}
		}
	}
	return somme
}

// Version optimisée du calcul de somme
func calculerSommeOptimise(tableau []int) int {
	somme := 0
	for _, val := range tableau {
		somme += val
	}
	return somme
}

// 2. Mesure de mémoire
func AnalyseMemoire() {
	// TIP: Utilisez Memory Profiler pour analyser l'utilisation de la mémoire
	// (Run > Profile 'nom_fonction' with Memory Profiler)

	// Création d'une grande quantité de données en mémoire
	var tableaux [][]int

	for i := 0; i < 100; i++ {
		// Allocation d'un grand tableau
		tableau := make([]int, 100000)
		for j := 0; j < 100000; j++ {
			tableau[j] = j
		}
		tableaux = append(tableaux, tableau)
	}

	// Utilisation des données pour qu'elles ne soient pas optimisées
	sommeTableaux := 0
	for _, tableau := range tableaux {
		sommeTableaux += len(tableau)
	}
	fmt.Printf("Total d'éléments alloués: %d\n", sommeTableaux)
}

// 3. Analyse de fuite de mémoire
func SimulerFuiteMemoire() {
	// TIP: Les fuites de mémoire peuvent être détectées avec le Memory Profiler

	// Simulation d'une fuite mémoire (conservation de références inutiles)
	cacheGlobal = make(map[string][]int)

	for i := 0; i < 1000; i++ {
		cle := fmt.Sprintf("data_%d", i)

		// Ajout de grandes quantités de données au cache sans jamais les nettoyer
		donnees := make([]int, 10000)
		for j := 0; j < 10000; j++ {
			donnees[j] = j
		}

		// Ajout au cache global sans mécanisme de nettoyage
		cacheGlobal[cle] = donnees
	}

	// Utilisation minimale pour éviter l'optimisation
	fmt.Printf("Taille du cache: %d entrées\n", len(cacheGlobal))
}

// Cache global (pour simuler une fuite mémoire)
var cacheGlobal map[string][]int

// 4. Optimisation basée sur le profiling
func ComparePerformance() {
	// TIP: Comparez les performances avec le CPU Profiler

	tableau := genererGrandTableau(100000)
	tableau2 := make([]int, len(tableau))
	copy(tableau2, tableau)

	// Mesure du temps pour la version inefficace
	debut := time.Now()
	trierTableauInefficace(tableau)
	dureeInefficace := time.Since(debut)

	// Mesure du temps pour la version optimisée
	debut = time.Now()
	trierTableauOptimise(tableau2)
	dureeOptimisee := time.Since(debut)

	fmt.Printf("Tri inefficace: %v\n", dureeInefficace)
	fmt.Printf("Tri optimisé: %v\n", dureeOptimisee)
	fmt.Printf("Ratio d'amélioration: %.2fx plus rapide\n",
		float64(dureeInefficace.Nanoseconds())/float64(dureeOptimisee.Nanoseconds()))
}

// ---------------------------------------------------------------------
// FONCTION DE DÉMONSTRATION
// ---------------------------------------------------------------------

// DemonstrationDebuggingProfiling exécute les différentes démos
func DemonstrationDebuggingProfiling() {
	fmt.Println("\n=== DÉMONSTRATION DU DÉBOGUEUR ===")
	fmt.Println("1. Points d'arrêt et inspection:")
	resultat := CalculationAvecBug(5, 10)
	fmt.Printf("Résultat de CalculationAvecBug: %d\n", resultat)

	fmt.Println("\n2. Évaluation d'expressions:")
	tableau := []int{1, 5, 10, 15, 20, 25, 30}
	somme := CalculerSommeTableau(tableau)
	fmt.Printf("Somme du tableau: %d\n", somme)

	fmt.Println("\n3. Points d'arrêt conditionnels:")
	tableau = []int{3, 15, 7, 12, 20, 9, 18, 25, 30, 11}
	resultatTableau := ProcessArray(tableau)
	fmt.Printf("Tableau traité: %v\n", resultatTableau)

	fmt.Println("\n4. Débogage de goroutines:")
	DeboguageGoroutines()

	// Note: Pour le profiling, il est préférable de lancer séparément chaque fonction
	// avec l'outil de profiling approprié (Run > Profile avec CPU ou Memory Profiler)
	fmt.Println("\n=== DÉMONSTRATION DU PROFILER ===")
	fmt.Println("Pour les démonstrations de profiling, utilisez:")
	fmt.Println("- Run > Profile 'FonctionLente' with CPU Profiler")
	fmt.Println("- Run > Profile 'AnalyseMemoire' with Memory Profiler")
	fmt.Println("- Run > Profile 'SimulerFuiteMemoire' with Memory Profiler")

	fmt.Println("\n5. Comparaison de performance:")
	ComparePerformance()
}
