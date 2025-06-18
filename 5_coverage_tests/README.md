# Démonstration des Fonctionnalités GoLand de JetBrains

Ce projet contient plusieurs démonstrations des fonctionnalités avancées de GoLand.

## 5. Couverture de Code et Tests

Cette section démontre comment GoLand gère les tests et l'analyse de couverture de code en Go.

### Fonctionnalités démontrées

- Exécution de tests unitaires
- Analyse de couverture de code
- Tests table-driven
- Navigation entre code et tests
- Génération automatique de tests
- Visualisation des résultats de tests

### Comment utiliser cette démonstration

1. **Exécuter tous les tests**:
   - Ouvrez le fichier `5_coverage_tests/coverage_tests_test.go`
   - Cliquez avec le bouton droit sur l'éditeur et sélectionnez "Run 'go test coverage_tests'"
   - Ou utilisez le raccourci clavier pour exécuter les tests (généralement Ctrl+Shift+F10)

2. **Analyse de couverture**:
   - Cliquez avec le bouton droit sur l'éditeur et sélectionnez "Run 'go test coverage_tests' with Coverage"
   - Observez les marqueurs de couverture dans le gutter (marge) de l'éditeur
   - Vérifiez le rapport de couverture dans la fenêtre "Coverage" en bas de l'IDE

3. **Exécuter un test spécifique**:
   - Placez le curseur à l'intérieur d'une fonction de test
   - Cliquez avec le bouton droit et sélectionnez "Run 'TestCalculerSomme'" (ou le nom du test)

4. **Navigation rapide entre le code et les tests**:
   - Dans un fichier de code, utilisez Ctrl+Shift+T pour basculer vers le fichier de test correspondant
   - Dans un fichier de test, utilisez la même combinaison pour revenir au code testé

5. **Génération de nouveaux tests**:
   - Ouvrez le fichier `5_coverage_tests/coverage_tests.go`
   - Placez le curseur sur une fonction sans test
   - Utilisez Alt+Enter et sélectionnez "Generate Test for function"

6. **Identification des zones non couvertes**:
   - Après avoir exécuté les tests avec couverture, les lignes non couvertes sont marquées en rouge
   - Notez que la fonction `AnalyserNote` a une couverture délibérément partielle

### Points à observer

- La visualisation de la couverture de code avec code coloré:
  - Vert: Code couvert par les tests
  - Rouge: Code non couvert
  - Jaune: Code partiellement couvert (conditions non entièrement testées)

- Les statistiques de couverture par package et par fonction

- L'intégration des tests dans l'environnement de développement

- La facilité de création et d'exécution des tests depuis l'IDE
