package ports

import "github.com/jackparsonss/grpc-monorepo/inventory/internal/application/domain"

type APIPort interface {
	AddInventoryItem(item *domain.InventoryItem) (*domain.InventoryItem, error)
	GetInventoryItem(id int64) (domain.InventoryItem, error)
}
