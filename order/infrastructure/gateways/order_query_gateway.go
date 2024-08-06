package gateways

import (
	"katsan"
	"katsan/order/application/repositories"
	"katsan/order/domain"

	accountQueries "katsan/account/application/queries"

	"github.com/jmoiron/sqlx"
)

type OrderQueryGateway struct {
	db *sqlx.DB
}

func NewOrderQueryGateway() repositories.OrderQueryRepository {
	return &OrderQueryGateway{
		db: katsan.NewDB(nil).GetDB(),
	}
}

func (g *OrderQueryGateway) LoadActiveOrders() ([]domain.Order, error) {
	orders := []domain.Order{}
	err := g.db.Select(&orders, queries.LoadActiveOrders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (g *OrderQueryGateway) LoadAccounts(pageSize, pageOffset string) ([]domain.Account, error) {
	accounts := []domain.Account{}

	accountServiceResult, accountServiceErr := accountQueries.NewAllActiveAccountsQueryHandler().Handle(accountQueries.AllActiveAccountsQuery{
		PageSize:   pageSize,
		PageOffset: pageOffset,
	})
	if accountServiceErr != nil {
		return nil, accountServiceErr
	}
	for _, result := range accountServiceResult {
		account := domain.Account{
			ID:   result.AccountID,
			Name: result.AccountName,
			Code: result.AccountCode,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (g *OrderQueryGateway) LoadAllOrdersByDay(date int) ([]domain.Order, error) {
	response := []domain.Order{}

	return response, nil
}
