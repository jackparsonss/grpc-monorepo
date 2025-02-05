package api

import (
	"context"
	"fmt"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/domain"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/ports"
)

type Application struct {
	inventory ports.InventoryPort
}

func NewApplication(inventory ports.InventoryPort) *Application {
	return &Application{
		inventory: inventory,
	}
}

func (a *Application) AddNewStock(item *domain.InventoryItem) (int64, error) {
	id, err := a.inventory.AddNewStock(context.Background(), item)
	if err != nil {
		fmt.Println("OG", err)
	}

	return id, nil
}

func (a *Application) GetStockItem(id int64) (*domain.InventoryItem, error) {
	item, err := a.inventory.GetStockItem(context.Background(), id)
	if err != nil {
		fmt.Println("OG", err)
		return nil, err
	}

	return item, nil
}
