package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// LongRunningProcess crée un processus qui s'exécute pendant un certain temps
// pour permettre de s'y attacher avec le débogueur.
func LongRunningProcess() {
	fmt.Println("Démarrage du processus longue durée...")
	fmt.Println("PID:", os.Getpid())
	fmt.Println("Vous pouvez maintenant vous attacher au processus avec GoLand:")
	fmt.Println("Run > Attach to Process > Sélectionnez ce processus")
	fmt.Println("-----------------------------------------------------")

	// Boucle infinie avec travail simulé
	counter := 0
	for {
		// Simule du travail
		processItem(counter)

		counter++

		// Ajoute une pause pour ralentir et permettre l'attachment
		time.Sleep(2 * time.Second)

		// Affiche une information pour montrer que le processus est toujours en cours
		fmt.Printf("Traitement en cours... Itération %d\n", counter)
	}
}

// processItem simule le traitement d'un élément
// Vous pouvez placer des points d'arrêt ici pour tester le débogueur
func processItem(id int) {
	// Génère un nombre aléatoire
	randomValue := rand.Intn(1000)

	// Effectue un calcul qui prend du temps
	result := 0
	for i := 0; i < 100000+randomValue; i++ {
		result += i % (id + 1)
	}

	// Affiche le résultat du traitement
	fmt.Printf("Item %d traité avec valeur aléatoire %d. Résultat: %d\n",
		id, randomValue, result)
}
