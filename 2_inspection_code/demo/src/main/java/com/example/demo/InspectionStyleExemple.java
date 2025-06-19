package com.example.demo;

import java.util.ArrayList;

/**
 * Démonstration des Inspections de Style de Code dans JetBrains IDE
 * Les IDE JetBrains détectent les problèmes de style et proposent des améliorations.
 */
public class InspectionStyleExemple {

    // 1. Nommage non conventionnel - L'IDE suggérera de suivre les conventions Java
    private String NomNonConventionnel;
    private int MAUVAIS_nom_variable;

    // 2. Méthode trop longue - L'IDE peut suggérer de la décomposer
    public void methodeTresLongue() {
        System.out.println("Début de la méthode");
        // Imaginons beaucoup de code ici...
        for (int i = 0; i < 100; i++) {
            if (i % 2 == 0) {
                System.out.println("Nombre pair: " + i);
            } else {
                System.out.println("Nombre impair: " + i);
            }

            // Plus de code...
            String texte = "Valeur actuelle: " + i;
            if (texte.length() > 10) {
                System.out.println("Texte long");
            } else {
                System.out.println("Texte court");
            }
        }
        System.out.println("Fin de la méthode");
    }

    // 3. Complexité cyclomatique élevée - L'IDE suggérera de simplifier
    public String complexiteElevee(int valeur) {
        if (valeur < 0) {
            return "Négatif";
        } else if (valeur == 0) {
            return "Zéro";
        } else if (valeur < 10) {
            if (valeur % 2 == 0) {
                return "Petit pair";
            } else {
                return "Petit impair";
            }
        } else if (valeur < 100) {
            if (valeur % 2 == 0) {
                if (valeur % 5 == 0) {
                    return "Moyen pair multiple de 5";
                } else {
                    return "Moyen pair non multiple de 5";
                }
            } else {
                return "Moyen impair";
            }
        } else {
            return "Grand";
        }
    }

    // 4. Type générique raw - L'IDE suggérera d'utiliser des génériques
    public void typeRawNonRecommande() {
        // L'IDE signalera l'utilisation de raw type au lieu du type générique
        ArrayList listeNonTypee = new ArrayList();  // Raw type
        listeNonTypee.add("texte");
        listeNonTypee.add(123);  // Mélange de types qui peut causer des problèmes
    }

    // 5. Instruction switch sans default - L'IDE peut le signaler
    public String moisSansDefault(int mois) {
        switch (mois) {
            case 1: return "Janvier";
            case 2: return "Février";
            case 3: return "Mars";
            // L'IDE signalera l'absence de case default
        }
        return null; // Potentiel NullPointerException si utilisé sans vérification
    }

    // 6. Utilisation préférable des enhanced for loops
    public void boucleAncienStyle() {
        int[] nombres = {1, 2, 3, 4, 5};

        // L'IDE suggérera d'utiliser un enhanced for loop
        for (int i = 0; i < nombres.length; i++) {
            System.out.println(nombres[i]);
        }

        // Version préférée que l'IDE suggérera:
        // for (int nombre : nombres) {
        //     System.out.println(nombre);
        // }
    }

    public static void main(String[] args) {
        System.out.println("Exemples d'inspection de style de code dans JetBrains IDE");
    }
}
