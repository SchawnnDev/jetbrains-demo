package inspection

// Démonstration des Inspections de Sécurité dans GoLand
// GoLand inclut des analyses de sécurité qui peuvent détecter des vulnérabilités potentielles.

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

// 1. Exemple d'inspection de vulnérabilités d'injection SQL
func InjectionSQL(db *sql.DB, userInput string) {
	// GoLand détectera le risque d'injection SQL ici
	query := "SELECT * FROM utilisateurs WHERE nom = '" + userInput + "'"
	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
}

// Solution sécurisée utilisant des requêtes préparées
func InjectionSQLSecurisee(db *sql.DB, userInput string) {
	// GoLand reconnaîtra cette approche comme sécurisée
	query := "SELECT * FROM utilisateurs WHERE nom = ?"
	rows, err := db.Query(query, userInput)
	if err != nil {
		return
	}
	defer rows.Close()
}

// 2. Exemple d'inspection de vulnérabilités d'injection de commandes
func InjectionCommande(userInput string) {
	// GoLand détectera le risque d'injection de commandes
	cmd := exec.Command("find", userInput)
	output, _ := cmd.Output()
	fmt.Println(string(output))
}

// 3. Exemple d'inspection de vulnérabilités XSS (Cross-Site Scripting)
func VulnerabiliteXSS(w http.ResponseWriter, userInput string) {
	// GoLand détectera le risque XSS ici
	fmt.Fprintf(w, "<div>%s</div>", userInput)
}

// Solution sécurisée utilisant l'échappement HTML
func XSSSecurise(w http.ResponseWriter, userInput string) {
	// GoLand reconnaîtra cette approche comme sécurisée
	tmpl := template.Must(template.New("exemple").Parse("<div>{{.}}</div>"))
	tmpl.Execute(w, userInput) // template.HTML échappe automatiquement
}

// 4. Exemple d'inspection d'utilisation non sécurisée de cookies
func CookiesNonSecurises(w http.ResponseWriter) {
	// GoLand suggérera d'ajouter des attributs de sécurité
	cookie := http.Cookie{
		Name:  "session",
		Value: "abc123",
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
}

// Solution sécurisée
func CookiesSecurises(w http.ResponseWriter) {
	// GoLand reconnaîtra cette approche comme plus sécurisée
	cookie := http.Cookie{
		Name:     "session",
		Value:    "abc123",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
}

// 5. Exemple d'inspection de gestion non sécurisée des mots de passe
func MotsDePasseNonSecurises() {
	// GoLand détectera l'utilisation de mots de passe en clair
	motDePasse := "password123"
	fmt.Println("Connexion avec le mot de passe:", motDePasse)

	// GoLand peut également détecter les mots de passe codés en dur
	connexionBDD("user", "password123")
}

func connexionBDD(utilisateur, motDePasse string) {
	// Fonction fictive pour l'exemple
}
