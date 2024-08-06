package domain

type Order struct {
	ID              string
	PlaceOfDelivery string
	DocumentNo      string
	DeliveryDate    string
	Status          string
	Items           []OrderItem
	AccountID       string
}

func (o *Order) TotalAmount() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Quantity
	}
	return total
}
