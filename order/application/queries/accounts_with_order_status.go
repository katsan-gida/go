package queries

import (
	"katsan/order/application/repositories"
	"katsan/order/infrastructure/gateways"
)

type GetDomesticAccountsWithOrderStatusQuery struct {
	PageOffset string
	PageSize   string
}

type GetDomesticAccountsWithOrderStatusReponse struct {
	AccountID   string
	AccountName string
	AccountCode string
	OrderStatus bool
}

type GetDomesticAccountsWithOrderStatusQueryHandler interface {
	Handle(query GetDomesticAccountsWithOrderStatusQuery) ([]GetDomesticAccountsWithOrderStatusReponse, error)
}

type getDomesticAccountsWithOrderStatusQueryHandler struct {
	repo repositories.OrderQueryRepository
}

func (h *getDomesticAccountsWithOrderStatusQueryHandler) Handle(query GetDomesticAccountsWithOrderStatusQuery) ([]GetDomesticAccountsWithOrderStatusReponse, error) {
	response := []GetDomesticAccountsWithOrderStatusReponse{}

	accounts, accountsErr := h.repo.LoadAccounts(query.PageSize, query.PageOffset)
	if accountsErr != nil {
		return nil, accountsErr
	}

	for _, account := range accounts {
		response = append(response, GetDomesticAccountsWithOrderStatusReponse{
			AccountID:   account.ID,
			AccountName: account.Name,
			AccountCode: account.Code,
		})
	}
	return response, nil
}

func NewGetDomesticAccountsWithOrderStatusQueryHandler() GetDomesticAccountsWithOrderStatusQueryHandler {
	return &getDomesticAccountsWithOrderStatusQueryHandler{
		repo: gateways.NewOrderQueryGateway(),
	}
}
