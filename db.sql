package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type Product struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Stock    int     `json:"stock"`
}

type Order struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Total     float64   `json:"total"`
    CreatedAt string    `json:"created_at"`
    Items     []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
    ID        int     `json:"id"`
    OrderID   int     `json:"order_id"`
    ProductID int     `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}
