package database_integration

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Driver PostgreSQL
)

// DB est l'instance de connexion à la base de données
var DB *sqlx.DB

// InitDB initialise la connexion à la base de données
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return fmt.Errorf("erreur de connexion à la base de données: %w", err)
	}

	// Configuration de la connexion
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Test de la connexion
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("erreur de ping à la base de données: %w", err)
	}

	log.Println("Connexion à la base de données établie avec succès")
	return nil
}

// CloseDB ferme la connexion à la base de données
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// UserRepository fournit des méthodes pour interagir avec les utilisateurs
type UserRepository struct{}

// GetUserByID récupère un utilisateur par son ID
func (r *UserRepository) GetUserByID(id int64) (*User, error) {
	user := &User{}
	err := DB.Get(user, "SELECT * FROM users WHERE user_id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Utilisateur non trouvé
		}
		return nil, err
	}
	return user, nil
}

// GetUserByUsername récupère un utilisateur par son nom d'utilisateur
func (r *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DB.Get(user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// CreateUser crée un nouvel utilisateur
func (r *UserRepository) CreateUser(user *User) error {
	query := `
		INSERT INTO users (username, email, first_name, last_name, password_hash, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING user_id, created_at`

	return DB.QueryRow(
		query,
		user.Username,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Password,
		user.IsActive,
	).Scan(&user.ID, &user.CreatedAt)
}

// UpdateUserLoginTime met à jour le moment de la dernière connexion
func (r *UserRepository) UpdateUserLoginTime(id int64) error {
	now := time.Now()
	_, err := DB.Exec(
		"UPDATE users SET last_login = $1 WHERE user_id = $2",
		now, id,
	)
	return err
}

// ProductRepository fournit des méthodes pour interagir avec les produits
type ProductRepository struct{}

// GetProductByID récupère un produit par son ID
func (r *ProductRepository) GetProductByID(id int64) (*Product, error) {
	product := &Product{}
	err := DB.Get(product, "SELECT * FROM products WHERE product_id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return product, nil
}

// GetProductsWithCategory récupère tous les produits avec leur catégorie
func (r *ProductRepository) GetProductsWithCategory() ([]*Product, error) {
	query := `
		SELECT p.*, c.name as category_name
		FROM products p
		JOIN categories c ON p.category_id = c.category_id
		ORDER BY p.product_id`

	rows, err := DB.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	for rows.Next() {
		var p struct {
			Product
			CategoryName string `db:"category_name"`
		}
		if err := rows.StructScan(&p); err != nil {
			return nil, err
		}

		p.Product.Category = &Category{
			ID:   p.CategoryID,
			Name: p.CategoryName,
		}
		products = append(products, &p.Product)
	}

	return products, nil
}

// CreateProduct crée un nouveau produit
func (r *ProductRepository) CreateProduct(product *Product) error {
	query := `
		INSERT INTO products (name, description, price, stock_quantity, category_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING product_id, created_at`

	return DB.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.StockQuantity,
		product.CategoryID,
	).Scan(&product.ID, &product.CreatedAt)
}

// UpdateProductStock met à jour le stock d'un produit
func (r *ProductRepository) UpdateProductStock(id int64, quantity int) error {
	_, err := DB.Exec(
		"UPDATE products SET stock_quantity = $1, updated_at = CURRENT_TIMESTAMP WHERE product_id = $2",
		quantity, id,
	)
	return err
}

// GetProductInventory récupère l'inventaire des produits (vue SQL)
func (r *ProductRepository) GetProductInventory() ([]*ProductInventory, error) {
	inventory := []*ProductInventory{}
	err := DB.Select(&inventory, "SELECT * FROM product_inventory ORDER BY product_id")
	return inventory, err
}

// OrderRepository fournit des méthodes pour interagir avec les commandes
type OrderRepository struct{}

// GetOrderByID récupère une commande par son ID
func (r *OrderRepository) GetOrderByID(id int64) (*Order, error) {
	order := &Order{}
	err := DB.Get(order, "SELECT * FROM orders WHERE order_id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// Récupérer les articles de la commande
	items := []*OrderItem{}
	err = DB.Select(&items, "SELECT * FROM order_items WHERE order_id = $1", id)
	if err != nil {
		return nil, err
	}
	order.OrderItems = items

	return order, nil
}

// CreateOrder crée une nouvelle commande avec ses articles
func (r *OrderRepository) CreateOrder(order *Order, items []*OrderItem) error {
	// Débuter une transaction
	tx, err := DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insérer la commande
	orderQuery := `
		INSERT INTO orders (user_id, status, shipping_address, billing_address, total_amount)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING order_id, order_date`

	err = tx.QueryRow(
		orderQuery,
		order.UserID,
		order.Status,
		order.ShippingAddress,
		order.BillingAddress,
		order.TotalAmount,
	).Scan(&order.ID, &order.OrderDate)

	if err != nil {
		return err
	}

	// Insérer les articles de la commande
	itemQuery := `
		INSERT INTO order_items (order_id, product_id, quantity, unit_price)
		VALUES ($1, $2, $3, $4)
		RETURNING order_item_id`

	for _, item := range items {
		item.OrderID = order.ID
		err = tx.QueryRow(
			itemQuery,
			item.OrderID,
			item.ProductID,
			item.Quantity,
			item.UnitPrice,
		).Scan(&item.ID)

		if err != nil {
			return err
		}
	}

	// Valider la transaction
	return tx.Commit()
}

// UpdateOrderStatus met à jour le statut d'une commande
func (r *OrderRepository) UpdateOrderStatus(id int64, status string) error {
	_, err := DB.Exec(
		"UPDATE orders SET status = $1 WHERE order_id = $2",
		status, id,
	)
	return err
}

// ProcessOrder traite une commande en utilisant la fonction SQL process_order
func (r *OrderRepository) ProcessOrder(id int64) (bool, error) {
	var success bool
	err := DB.Get(&success, "SELECT process_order($1)", id)
	return success, err
}

// GetOrderDetails récupère les détails d'une commande via la vue SQL
func (r *OrderRepository) GetOrderDetails() ([]*OrderDetails, error) {
	details := []*OrderDetails{}
	err := DB.Select(&details, "SELECT * FROM order_details ORDER BY order_date DESC")
	return details, err
}
