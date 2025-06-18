package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Récupération du port depuis les variables d'environnement
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut
	}

	// Définition du handler pour la racine
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requête reçue : %s %s", r.Method, r.URL.Path)
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "inconnu"
		}
		fmt.Fprintf(w, "Bonjour depuis GoLand et Docker! 🐳\n\n")
		fmt.Fprintf(w, "Serveur: %s\n", hostname)
		fmt.Fprintf(w, "Version Go: %s\n", getGoVersion())
		fmt.Fprintf(w, "Environnement:\n")

		// Affichage des variables d'environnement
		for _, env := range os.Environ() {
			fmt.Fprintf(w, "  %s\n", env)
		}
	})

	// Démarrage du serveur
	serverAddr := ":" + port
	log.Printf("Démarrage du serveur sur %s...", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}

// getGoVersion retourne la version de Go utilisée
func getGoVersion() string {
	return "go1.21" // À remplacer par une détection automatique au besoin
}
