package repositories

import "katsan/account/domain"

type AccountQueryRepository interface {
	LoadAllActiveAccounts(pageSize, pageOffset string) ([]domain.Account, error)
}
