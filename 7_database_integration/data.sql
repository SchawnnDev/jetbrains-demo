-- Fichier data.sql : Insertion de données exemple pour démontrer l'intégration de base de données dans JetBrains

-- Insertion des données utilisateur
INSERT INTO users (username, email, first_name, last_name, password_hash, is_active)
VALUES
    ('admin', 'admin@example.com', 'Admin', 'User', '$2a$12$tH8S.qNFEwf0zQhPeF9hG.3vrYhGqEDKI7LaaNgDLTIfeKYEcHE8K', TRUE),
    ('jdupont', 'jean.dupont@example.com', 'Jean', 'Dupont', '$2a$12$4I85imG5BW/n.HnEIS7r9uutvVKGi2HhXYTOgOK06MBQ2p9QqcOJu', TRUE),
    ('mleblanc', 'marie.leblanc@example.com', 'Marie', 'Leblanc', '$2a$12$YETTW0DYUNOpli1Cf9Ef/.odMyPZ2Q2EEg3M0SSA.YZlB2bXaOKzW', TRUE),
    ('pmartin', 'pierre.martin@example.com', 'Pierre', 'Martin', '$2a$12$FRWAGowt/Ji/5YwRG2m5PuvZLjZKPgqkBdoCI.znJkxZtQiJLkqgO', FALSE),
    ('srobert', 'sophie.robert@example.com', 'Sophie', 'Robert', '$2a$12$KmMJSH6gGSbP6XDbMoXkLeyVEYnEVQGObIjkuIpPmCIl7WV9BOGfK', TRUE);

-- Insertion des catégories
INSERT INTO categories (name, description, parent_id)
VALUES
    ('Électronique', 'Produits électroniques et appareils', NULL),
    ('Ordinateurs', 'Ordinateurs et accessoires', 1),
    ('Smartphones', 'Téléphones intelligents et accessoires', 1),
    ('Audio', 'Équipement audio', 1),
    ('Vêtements', 'Vêtements et accessoires de mode', NULL),
    ('Homme', 'Vêtements pour homme', 5),
    ('Femme', 'Vêtements pour femme', 5),
    ('Enfant', 'Vêtements pour enfant', 5),
    ('Livres', 'Livres et publications', NULL),
    ('Fiction', 'Romans et histoires de fiction', 9);

-- Insertion des produits
INSERT INTO products (name, description, price, stock_quantity, category_id)
VALUES
    ('Laptop Pro X', 'Ordinateur portable haut de gamme avec processeur i9', 1299.99, 15, 2),
    ('Smartphone Galaxy', 'Smartphone dernière génération avec écran AMOLED', 899.99, 25, 3),
    ('Écouteurs sans fil', 'Écouteurs Bluetooth avec réduction de bruit active', 199.99, 50, 4),
    ('Chemise formelle', 'Chemise en coton pour occasions spéciales', 59.99, 30, 6),
    ('Robe d''été', 'Robe légère et élégante pour l''été', 79.99, 20, 7),
    ('T-shirt enfant', 'T-shirt coloré pour enfant avec motifs', 19.99, 45, 8),
    ('Le Guide du développeur Go', 'Livre pour apprendre le langage Go', 39.99, 10, 9),
    ('Moniteur 4K', 'Écran 4K ultra HD 32 pouces', 399.99, 8, 2),
    ('Tablette MediaPad', 'Tablette 10 pouces avec grand espace de stockage', 349.99, 12, 2),
    ('Enceinte Bluetooth', 'Enceinte portable étanche', 89.99, 18, 4);

-- Insertion des commandes
INSERT INTO orders (user_id, status, shipping_address, billing_address, total_amount)
VALUES
    (2, 'delivered', '15 Rue des Fleurs, 75001 Paris', '15 Rue des Fleurs, 75001 Paris', 1499.98),
    (3, 'shipped', '27 Avenue Victor Hugo, 69002 Lyon', '27 Avenue Victor Hugo, 69002 Lyon', 979.98),
    (4, 'pending', '8 Boulevard Jean Jaurès, 33000 Bordeaux', '8 Boulevard Jean Jaurès, 33000 Bordeaux', 419.97),
    (5, 'processing', '42 Rue de la Paix, 59000 Lille', '42 Rue de la Paix, 59000 Lille', 139.98),
    (2, 'cancelled', '15 Rue des Fleurs, 75001 Paris', '15 Rue des Fleurs, 75001 Paris', 399.99);

-- Insertion des articles commandés
INSERT INTO order_items (order_id, product_id, quantity, unit_price)
VALUES
    (1, 1, 1, 1299.99),
    (1, 3, 1, 199.99),
    (2, 2, 1, 899.99),
    (2, 4, 1, 79.99),
    (3, 8, 1, 399.99),
    (3, 6, 1, 19.98),
    (4, 7, 1, 39.99),
    (4, 6, 5, 19.99),
    (5, 8, 1, 399.99);
