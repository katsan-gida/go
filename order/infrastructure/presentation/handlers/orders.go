package handlers

import (
	"katsan/order/infrastructure/presentation/pages"
	"net/http"
)

func HandleGetOrdersPage(w http.ResponseWriter, r *http.Request) {
	pages.OrdersPage().Render(r.Context(), w)
}

func HandleGetNewOrderPage(w http.ResponseWriter, r *http.Request) {
	pages.NewOrderPage().Render(r.Context(), w)
}
