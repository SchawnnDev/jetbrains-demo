# Utilisation d'une image multi-stage pour optimiser la taille finale
# Étape 1: Build de l'application
FROM golang:1.24-alpine AS builder

# Installation des dépendances système nécessaires
RUN apk add --no-cache git

# Définition du répertoire de travail
WORKDIR /app

# Copie des fichiers Go et téléchargement des dépendances
COPY go.mod ./
# Si vous avez un fichier go.sum, décommentez la ligne suivante
# COPY go.sum ./
RUN go mod download

# Copie du code source
COPY 6_docker_integration/ ./

# Compilation du programme avec optimisations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o goapp app.go

# Étape 2: Création de l'image finale légère
FROM alpine:latest

# Installation des certificats pour les connexions HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Définition des variables d'environnement
ENV PORT=8080
ENV GO_ENV=production

# Définition du répertoire de travail
WORKDIR /app

# Copie du binaire compilé depuis l'étape précédente
COPY --from=builder /app/goapp .

# Exposition du port utilisé par l'application
EXPOSE $PORT

# Commande de démarrage de l'application
CMD ["./goapp"]
