package ports

import "github.com/jackparsonss/grpc-monorepo/inventory/internal/application/domain"

type DBPort interface {
	GetInventoryItem(id int64) (domain.InventoryItem, error)
	CreateInventoryItem(item *domain.InventoryItem) error
}
