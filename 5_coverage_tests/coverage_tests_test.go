package coverage_tests

import (
	"testing"
)

// ---------------------------------------------------------------------
// Tests unitaires pour les fonctions simples
// ---------------------------------------------------------------------

// Test de CalculerSomme
func TestCalculerSomme(t *testing.T) {
	// Structure de test table-driven
	tests := []struct {
		nom     string
		a, b    int
		attendu int
	}{
		{"Somme de positifs", 5, 3, 8},
		{"Somme avec zéro", 5, 0, 5},
		{"Somme de négatifs", -2, -3, -5},
		{"Somme positif et négatif", 5, -3, 2},
	}

	// Exécution des cas de test
	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := CalculerSomme(test.a, test.b)
			if resultat != test.attendu {
				t.Errorf("CalculerSomme(%d, %d) = %d; attendu %d",
					test.a, test.b, resultat, test.attendu)
			}
		})
	}
}

// Test de CalculerProduit
func TestCalculerProduit(t *testing.T) {
	// Structure de test table-driven
	tests := []struct {
		nom     string
		a, b    int
		attendu int
	}{
		{"Produit de positifs", 5, 3, 15},
		{"Produit avec zéro", 5, 0, 0},
		{"Produit de négatifs", -2, -3, 6},
		{"Produit positif et négatif", 5, -3, -15},
	}

	// Exécution des cas de test
	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := CalculerProduit(test.a, test.b)
			if resultat != test.attendu {
				t.Errorf("CalculerProduit(%d, %d) = %d; attendu %d",
					test.a, test.b, resultat, test.attendu)
			}
		})
	}
}

// Test de CalculerQuotient
func TestCalculerQuotient(t *testing.T) {
	// Test de cas normaux
	t.Run("Division valide", func(t *testing.T) {
		resultat, err := CalculerQuotient(10, 2)
		if err != nil {
			t.Errorf("CalculerQuotient(10, 2) a retourné une erreur: %v", err)
		}
		if resultat != 5.0 {
			t.Errorf("CalculerQuotient(10, 2) = %f; attendu 5.0", resultat)
		}
	})

	// Test de division par zéro
	t.Run("Division par zéro", func(t *testing.T) {
		_, err := CalculerQuotient(10, 0)
		if err == nil {
			t.Error("CalculerQuotient(10, 0) devrait retourner une erreur")
		}
	})
}

// Test de VerifierNombrePositif
func TestVerifierNombrePositif(t *testing.T) {
	tests := []struct {
		n       int
		attendu bool
	}{
		{5, true},
		{0, false},
		{-3, false},
	}

	for _, test := range tests {
		resultat := VerifierNombrePositif(test.n)
		if resultat != test.attendu {
			t.Errorf("VerifierNombrePositif(%d) = %v; attendu %v",
				test.n, resultat, test.attendu)
		}
	}
}

// Test de TraiterTexte
func TestTraiterTexte(t *testing.T) {
	tests := []struct {
		nom     string
		texte   string
		attendu string
	}{
		{"Texte avec espaces", "hello world", "HELLOWORLD"},
		{"Texte sans espaces", "hello", "HELLO"},
		{"Texte vide", "", ""},
		{"Texte avec caractères spéciaux", "Hello, World!", "HELLO,WORLD!"},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := TraiterTexte(test.texte)
			if resultat != test.attendu {
				t.Errorf("TraiterTexte(%q) = %q; attendu %q",
					test.texte, resultat, test.attendu)
			}
		})
	}
}

// ---------------------------------------------------------------------
// Tests unitaires pour la structure Calculatrice
// ---------------------------------------------------------------------

// Test des méthodes de Calculatrice
func TestCalculatrice(t *testing.T) {
	calc := Calculatrice{}

	// Test de l'état initial
	t.Run("Mémoire initiale", func(t *testing.T) {
		if calc.ObtenirMemoire() != 0 {
			t.Errorf("Mémoire initiale = %d; attendu 0", calc.ObtenirMemoire())
		}
	})

	// Test de la méthode Ajouter
	t.Run("Ajouter", func(t *testing.T) {
		calc.Ajouter(5)
		if calc.ObtenirMemoire() != 5 {
			t.Errorf("Après Ajouter(5) = %d; attendu 5", calc.ObtenirMemoire())
		}
	})

	// Test de la méthode Soustraire
	t.Run("Soustraire", func(t *testing.T) {
		calc.Soustraire(2)
		if calc.ObtenirMemoire() != 3 {
			t.Errorf("Après Soustraire(2) = %d; attendu 3", calc.ObtenirMemoire())
		}
	})

	// Test de la méthode ReinitialiserMemoire
	t.Run("Réinitialiser", func(t *testing.T) {
		calc.ReinitialiserMemoire()
		if calc.ObtenirMemoire() != 0 {
			t.Errorf("Après ReinitialiserMemoire() = %d; attendu 0", calc.ObtenirMemoire())
		}
	})
}

// ---------------------------------------------------------------------
// Tests partiels pour démontrer les analyses de couverture
// ---------------------------------------------------------------------

// Test de AnalyserNote - couverture partielle pour démonstration
func TestAnalyserNote(t *testing.T) {
	tests := []struct {
		note    int
		attendu string
	}{
		{95, "Excellent"},
		{85, "Très bien"},
		{-5, "Note invalide: nombre négatif"},
		{45, "Insuffisant"},
	}

	for _, test := range tests {
		resultat := AnalyserNote(test.note)
		if resultat != test.attendu {
			t.Errorf("AnalyserNote(%d) = %q; attendu %q",
				test.note, resultat, test.attendu)
		}
	}
	// Note: Ce test ne couvre pas toutes les branches de la fonction AnalyserNote
}

// Test de FiltrerTableau
func TestFiltrerTableau(t *testing.T) {
	tableau := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Fonction de filtre pour les nombres pairs
	estPair := func(n int) bool {
		return n%2 == 0
	}

	resultat := FiltrerTableau(tableau, estPair)
	attendu := []int{2, 4, 6, 8, 10}

	// Vérification de la taille du résultat
	if len(resultat) != len(attendu) {
		t.Errorf("FiltrerTableau() retourne %d éléments; attendu %d",
			len(resultat), len(attendu))
	}

	// Vérification des éléments du résultat
	for i, v := range resultat {
		if v != attendu[i] {
			t.Errorf("FiltrerTableau()[%d] = %d; attendu %d", i, v, attendu[i])
		}
	}
}
