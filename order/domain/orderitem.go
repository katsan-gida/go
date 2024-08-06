package domain

type OrderItem struct {
	OrderItemID string
	OrderID     string
	ProductID   string
	Quantity    float64
	Unit        string
}
