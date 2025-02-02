package domain

import "time"

type InventoryItem struct {
	ID          int64   `json:"id"`
	Price       float32 `json:"price"`
	Title       string  `json:"title"`
	Company     string  `json:"company"`
	Description string  `json:"description"`
	StockCount  int64   `json:"stock_count"`
	CreatedAt   int64   `json:"created_at"`
}

func NewInventoryItem(price float32, title, company, description string, stockCount int64) *InventoryItem {
	return &InventoryItem{
		Price:       price,
		Title:       title,
		Company:     company,
		Description: description,
		StockCount:  stockCount,
		CreatedAt:   time.Now().UnixNano(),
	}
}
