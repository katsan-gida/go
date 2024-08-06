package gateways

import (
	"katsan"
	"katsan/account/application/repositories"
	"katsan/account/domain"

	"github.com/jmoiron/sqlx"
)

type AccountQueryGateway struct {
	db *sqlx.DB
}

func NewAccountQueryGateway() repositories.AccountQueryRepository {
	return &AccountQueryGateway{
		db: katsan.NewDB(nil).GetDB(),
	}
}

func (g *AccountQueryGateway) LoadAllActiveAccounts(pageSize, pageOffset string) ([]domain.Account, error) {
	accounts := []domain.Account{}
	err := g.db.Select(&accounts, queries.LoadAllActiveAccounts, pageSize, pageOffset)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
