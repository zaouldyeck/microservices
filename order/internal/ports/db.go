package ports

import "github.com/zaouldyeck/microservices/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
