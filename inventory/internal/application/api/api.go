package api

import (
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/application/domain"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func (a Application) GetInventoryItem(id int64) (domain.InventoryItem, error) {
	return a.db.GetInventoryItem(id)
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) AddInventoryItem(item *domain.InventoryItem) (*domain.InventoryItem, error) {
	err := a.db.CreateInventoryItem(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
