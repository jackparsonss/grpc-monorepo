package domain

type InventoryItem struct {
	ID          int64   `json:"id,omitempty"`
	Price       float32 `json:"price"`
	Title       string  `json:"title"`
	Company     string  `json:"company"`
	Description string  `json:"description"`
	StockCount  int64   `json:"stock_count"`
}
