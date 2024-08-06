package queries

import (
	"katsan/account/application/repositories"
	"katsan/account/infrastructure/gateways"
)

type AllActiveAccountsQuery struct {
	PageSize   string
	PageOffset string
}

type AllActiveAccountsQueryResponse struct {
	AccountID   string
	AccountName string
	AccountCode string
}

type AllActiveAccountsQueryHandler interface {
	Handle(AllActiveAccountsQuery) ([]AllActiveAccountsQueryResponse, error)
}

type allActiveAccountsQueryHandler struct {
	repo repositories.AccountQueryRepository
}

func NewAllActiveAccountsQueryHandler() AllActiveAccountsQueryHandler {
	return &allActiveAccountsQueryHandler{repo: gateways.NewAccountQueryGateway()}
}

func (h *allActiveAccountsQueryHandler) Handle(query AllActiveAccountsQuery) ([]AllActiveAccountsQueryResponse, error) {
	response := []AllActiveAccountsQueryResponse{}

	accounts, err := h.repo.LoadAllActiveAccounts(query.PageSize, query.PageOffset)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		response = append(response, AllActiveAccountsQueryResponse{
			AccountID:   account.ID,
			AccountName: account.Name1 + " " + account.Name2,
			AccountCode: account.Code,
		})
	}

	return response, nil
}
