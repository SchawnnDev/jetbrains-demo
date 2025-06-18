package database_integration

import (
	"time"
)

// User représente un utilisateur du système
type User struct {
	ID        int64      `db:"user_id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Password  string     `db:"-"` // Non stocké directement, seul le hash est stocké
	IsActive  bool       `db:"is_active"`
	CreatedAt time.Time  `db:"created_at"`
	LastLogin *time.Time `db:"last_login"`
}

// Category représente une catégorie de produit
type Category struct {
	ID          int64     `db:"category_id"`
	Name        string    `db:"name"`
	Description *string   `db:"description"`
	ParentID    *int64    `db:"parent_id"`
	CreatedAt   time.Time `db:"created_at"`

	// Relations
	Parent   *Category   `db:"-"`
	Children []*Category `db:"-"`
}

// Product représente un produit dans le catalogue
type Product struct {
	ID            int64      `db:"product_id"`
	Name          string     `db:"name"`
	Description   *string    `db:"description"`
	Price         float64    `db:"price"`
	StockQuantity int        `db:"stock_quantity"`
	CategoryID    int64      `db:"category_id"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`

	// Relations
	Category *Category `db:"-"`
}

// Order représente une commande client
type Order struct {
	ID              int64     `db:"order_id"`
	UserID          int64     `db:"user_id"`
	OrderDate       time.Time `db:"order_date"`
	Status          string    `db:"status"`
	ShippingAddress string    `db:"shipping_address"`
	BillingAddress  string    `db:"billing_address"`
	TotalAmount     float64   `db:"total_amount"`

	// Relations
	User       *User        `db:"-"`
	OrderItems []*OrderItem `db:"-"`
}

// OrderItem représente un article dans une commande
type OrderItem struct {
	ID        int64   `db:"order_item_id"`
	OrderID   int64   `db:"order_id"`
	ProductID int64   `db:"product_id"`
	Quantity  int     `db:"quantity"`
	UnitPrice float64 `db:"unit_price"`

	// Relations
	Order   *Order   `db:"-"`
	Product *Product `db:"-"`
}

// ProductInventory représente une entrée de la vue product_inventory
type ProductInventory struct {
	ProductID      int64   `db:"product_id"`
	ProductName    string  `db:"product_name"`
	CategoryName   string  `db:"category_name"`
	Price          float64 `db:"price"`
	StockQuantity  int     `db:"stock_quantity"`
	InventoryValue float64 `db:"inventory_value"`
}

// OrderDetails représente une entrée de la vue order_details
type OrderDetails struct {
	OrderID     int64     `db:"order_id"`
	OrderDate   time.Time `db:"order_date"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	Status      string    `db:"status"`
	TotalAmount float64   `db:"total_amount"`
	TotalItems  int       `db:"total_items"`
}
