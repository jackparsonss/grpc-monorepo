package ports

import (
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/domain"
)

type APIPort interface {
	AddNewStock(item *domain.InventoryItem) (int64, error)
	GetStockItem(id int64) (*domain.InventoryItem, error)
}
