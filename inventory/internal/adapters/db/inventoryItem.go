package db

import (
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/application/domain"
	"gorm.io/gorm"
)

type InventoryItem struct {
	gorm.Model
	Price       float32 `json:"price"`
	Title       string  `json:"title"`
	Company     string  `json:"company"`
	Description string  `json:"description"`
	StockCount  int64   `json:"stock_count"`
}

func (a *Adapter) GetInventoryItem(id int64) (domain.InventoryItem, error) {
	var inventoryEntity InventoryItem
	res := a.db.First(&inventoryEntity, id)

	order := domain.InventoryItem{
		ID:          int64(inventoryEntity.ID),
		Title:       inventoryEntity.Title,
		Price:       inventoryEntity.Price,
		Company:     inventoryEntity.Company,
		Description: inventoryEntity.Description,
		StockCount:  inventoryEntity.StockCount,
		CreatedAt:   inventoryEntity.CreatedAt.UnixNano(),
	}

	return order, res.Error
}

func (a *Adapter) CreateInventoryItem(inventoryItem *domain.InventoryItem) error {
	inventoryItemEntity := InventoryItem{
		Title:       inventoryItem.Title,
		Price:       inventoryItem.Price,
		Company:     inventoryItem.Company,
		Description: inventoryItem.Description,
		StockCount:  inventoryItem.StockCount,
	}

	res := a.db.Create(&inventoryItemEntity)
	if res.Error != nil {
		return res.Error
	}

	inventoryItem.ID = int64(inventoryItemEntity.ID)
	return nil
}
