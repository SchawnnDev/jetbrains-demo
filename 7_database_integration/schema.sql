-- Fichier schema.sql : Définition du schéma de base de données pour l'exemple JetBrains Database

-- Création des tables
-- -----------------------------------------------------

-- Table: users
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);
COMMENT ON TABLE users IS 'Table contenant les informations utilisateurs';

-- Table: categories
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    parent_id INTEGER NULL REFERENCES categories(category_id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE categories IS 'Catégories de produits';

-- Table: products
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    stock_quantity INTEGER NOT NULL DEFAULT 0 CHECK (stock_quantity >= 0),
    category_id INTEGER NOT NULL REFERENCES categories(category_id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
COMMENT ON TABLE products IS 'Catalogue de produits';
CREATE INDEX idx_products_category ON products(category_id);

-- Table: orders
CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'processing', 'shipped', 'delivered', 'cancelled')),
    shipping_address TEXT NOT NULL,
    billing_address TEXT NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL CHECK (total_amount >= 0)
);
COMMENT ON TABLE orders IS 'Commandes clients';
CREATE INDEX idx_orders_user ON orders(user_id);

-- Table: order_items
CREATE TABLE order_items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(order_id) ON DELETE CASCADE,
    product_id INTEGER NOT NULL REFERENCES products(product_id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10, 2) NOT NULL CHECK (unit_price >= 0),
    UNIQUE (order_id, product_id)
);
COMMENT ON TABLE order_items IS 'Détails des articles commandés';
CREATE INDEX idx_order_items_order ON order_items(order_id);
CREATE INDEX idx_order_items_product ON order_items(product_id);

-- Création des vues
-- -----------------------------------------------------

-- Vue: product_inventory
CREATE VIEW product_inventory AS
SELECT
    p.product_id,
    p.name AS product_name,
    c.name AS category_name,
    p.price,
    p.stock_quantity,
    p.stock_quantity * p.price AS inventory_value
FROM products p
JOIN categories c ON p.category_id = c.category_id;
COMMENT ON VIEW product_inventory IS 'Vue des produits avec leur valeur d inventaire';

-- Vue: order_details
CREATE VIEW order_details AS
SELECT
    o.order_id,
    o.order_date,
    u.username,
    u.email,
    o.status,
    o.total_amount,
    COUNT(oi.order_item_id) AS total_items
FROM orders o
JOIN users u ON o.user_id = u.user_id
JOIN order_items oi ON o.order_id = oi.order_id
GROUP BY o.order_id, u.username, u.email;
COMMENT ON VIEW order_details IS 'Vue détaillée des commandes avec informations clients';

-- Création des fonctions et procédures
-- -----------------------------------------------------

-- Fonction: update_product_price
CREATE OR REPLACE FUNCTION update_product_price() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger pour mettre à jour la date de modification lors d'une mise à jour de prix
CREATE TRIGGER trg_update_product_price
BEFORE UPDATE OF price ON products
FOR EACH ROW
EXECUTE FUNCTION update_product_price();

-- Fonction: calculate_order_total
CREATE OR REPLACE FUNCTION calculate_order_total(order_id_param INTEGER) RETURNS DECIMAL AS $$
DECLARE
    total DECIMAL(10, 2);
BEGIN
    SELECT COALESCE(SUM(unit_price * quantity), 0) INTO total
    FROM order_items
    WHERE order_id = order_id_param;

    RETURN total;
END;
$$ LANGUAGE plpgsql;

-- Fonction: process_order
CREATE OR REPLACE FUNCTION process_order(order_id_param INTEGER) RETURNS BOOLEAN AS $$
DECLARE
    valid BOOLEAN := TRUE;
    item RECORD;
BEGIN
    -- Vérifier le stock pour chaque article
    FOR item IN
        SELECT oi.product_id, oi.quantity, p.stock_quantity
        FROM order_items oi
        JOIN products p ON oi.product_id = p.product_id
        WHERE oi.order_id = order_id_param
    LOOP
        -- Si pas assez de stock
        IF item.quantity > item.stock_quantity THEN
            valid := FALSE;
            EXIT;
        END IF;
    END LOOP;

    -- Si commande valide, mettre à jour le stock et statut de la commande
    IF valid THEN
        -- Mettre à jour le statut de la commande
        UPDATE orders SET status = 'processing' WHERE order_id = order_id_param;

        -- Mettre à jour le stock
        UPDATE products p
        SET stock_quantity = p.stock_quantity - oi.quantity
        FROM order_items oi
        WHERE oi.order_id = order_id_param AND p.product_id = oi.product_id;
    END IF;

    RETURN valid;
END;
$$ LANGUAGE plpgsql;
