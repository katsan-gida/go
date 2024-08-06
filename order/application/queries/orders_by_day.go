package queries

import (
	"katsan/order/application/repositories"
	"katsan/order/infrastructure/gateways"
)

type AllOrdersByDayQuery struct {
}

type AllOrdersByDayResponse struct {
}

type AllOrdersByDayQueryHandler interface {
	Handle(query AllOrdersByDayQuery) ([]AllOrdersByDayResponse, error)
}

type allOrdersByDayQueryHandler struct {
	repo repositories.OrderQueryRepository
}

func (h *allOrdersByDayQueryHandler) Handle(query AllOrdersByDayQuery) ([]AllOrdersByDayResponse, error) {
	response := []AllOrdersByDayResponse{}
	return response, nil
}

func NewAllOrdersByDayQueryHandler() AllOrdersByDayQueryHandler {
	return &allOrdersByDayQueryHandler{repo: gateways.NewOrderQueryGateway()}
}
