package ports

import (
	"context"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/domain"
)

type InventoryPort interface {
	AddNewStock(ctx context.Context, item *domain.InventoryItem) (int64, error)
	GetStockItem(ctx context.Context, id int64) (*domain.InventoryItem, error)
}
