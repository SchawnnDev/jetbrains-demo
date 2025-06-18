package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("=== Démonstrateur de Debugging et Profiling dans GoLand ===")

	// Vérification des arguments de la ligne de commande
	if len(os.Args) < 2 {
		afficherAide()
		return
	}

	// Conversion de l'argument en nombre
	choix, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erreur: Veuillez fournir un numéro de démo valide")
		afficherAide()
		return
	}

	// Exécution de la fonction sélectionnée
	switch choix {
	case 1:
		fmt.Println("\n=== Démo 1: Points d'arrêt et inspection des variables ===")
		resultat := CalculationAvecBug(5, 10)
		fmt.Printf("Résultat de CalculationAvecBug: %d\n", resultat)

	case 2:
		fmt.Println("\n=== Démo 2: Évaluation d'expressions pendant le débogage ===")
		tableau := []int{1, 5, 10, 15, 20, 25, 30}
		somme := CalculerSommeTableau(tableau)
		fmt.Printf("Somme du tableau: %d\n", somme)

	case 3:
		fmt.Println("\n=== Démo 3: Points d'arrêt conditionnels et watchpoints ===")
		tableau := []int{3, 15, 7, 12, 20, 9, 18, 25, 30, 11}
		resultatTableau := ProcessArray(tableau)
		fmt.Printf("Tableau traité: %v\n", resultatTableau)

	case 4:
		fmt.Println("\n=== Démo 4: Débogage de goroutines ===")
		DeboguageGoroutines()

	case 5:
		fmt.Println("\n=== Démo 5: Analyse de performances - Fonction lente ===")
		fmt.Println("Exécution de FonctionLente()...")
		fmt.Println("TIP: Lancez cette démo avec le CPU Profiler via Run > Profile...")
		FonctionLente()

	case 6:
		fmt.Println("\n=== Démo 6: Analyse de mémoire ===")
		fmt.Println("Exécution de AnalyseMemoire()...")
		fmt.Println("TIP: Lancez cette démo avec le Memory Profiler via Run > Profile...")
		AnalyseMemoire()

	case 7:
		fmt.Println("\n=== Démo 7: Simulation de fuite mémoire ===")
		fmt.Println("Exécution de SimulerFuiteMemoire()...")
		fmt.Println("TIP: Lancez cette démo avec le Memory Profiler via Run > Profile...")
		SimulerFuiteMemoire()

	case 8:
		fmt.Println("\n=== Démo 9: Attach to Process Debugger ===")
		fmt.Println("Lancement d'un processus de longue durée pour démonstration de Attach to Process...")
		LongRunningProcess()

	case 0:
		fmt.Println("\n=== Exécution de toutes les démos ===")
		DemonstrationDebuggingProfiling()

	default:
		fmt.Println("Erreur: Choix invalide")
		afficherAide()
	}
}

func afficherAide() {
	fmt.Println("\nUtilisation: go run main.go <numéro_de_démo>")
	fmt.Println("\nDémos disponibles:")
	fmt.Println("  1: Points d'arrêt et inspection des variables")
	fmt.Println("  2: Évaluation d'expressions pendant le débogage")
	fmt.Println("  3: Points d'arrêt conditionnels et watchpoints")
	fmt.Println("  4: Débogage de goroutines")
	fmt.Println("  5: Analyse de performances - Fonction lente (à utiliser avec CPU Profiler)")
	fmt.Println("  6: Analyse de mémoire (à utiliser avec Memory Profiler)")
	fmt.Println("  7: Simulation de fuite mémoire (à utiliser avec Memory Profiler)")
	fmt.Println("  8: Attacher au débogueur de processus")
	fmt.Println("  0: Exécuter toutes les démos")
	fmt.Println("\nExemple: go run main.go 1")
	fmt.Println("\nTIP: Pour utiliser le profiler, sélectionnez la configuration de démo correspondante,")
	fmt.Println("     puis utilisez Run > Profile... et choisissez CPU ou Memory Profiler.")
}
