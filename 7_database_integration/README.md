# Intégration de Base de Données avec JetBrains

Cette partie du projet démontre les fonctionnalités avancées d'intégration de bases de données dans les outils JetBrains, notamment avec PostgreSQL et Go.

## Configuration de la Base de Données PostgreSQL

### Démarrer PostgreSQL avec Docker

La méthode la plus simple pour démarrer rapidement une base de données PostgreSQL à des fins de test est d'utiliser Docker.

```bash
# Démarrer un conteneur PostgreSQL
docker run --name goland-db-demo -e POSTGRES_PASSWORD=jetbrains_password -e POSTGRES_USER=jetbrains -e POSTGRES_DB=jetbrains_demo -p 5432:5432 -d postgres:16-alpine
```


