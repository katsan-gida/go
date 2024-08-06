package repositories

import "katsan/order/domain"

type OrderQueryRepository interface {
	LoadActiveOrders() ([]domain.Order, error)
	LoadAccounts(pageSize, pageOffset string) ([]domain.Account, error)
	LoadAllOrdersByDay(date int) ([]domain.Order, error)
}
