package inventory

import (
	"context"

	"github.com/jackparsonss/grpc-monorepo/proto/golang/inventory"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	inventory inventory.InventoryClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := inventory.NewInventoryClient(conn)
	return &Adapter{inventory: client}, nil
}

func (a *Adapter) AddNewStock(ctx context.Context, item *domain.InventoryItem) (int64, error) {
	grpcItem := &inventory.InventoryItem{
		Price:       item.Price,
		Company:     item.Company,
		Description: item.Description,
		StockCount:  item.StockCount,
		Title:       item.Title,
	}

	res, err := a.inventory.CreateItem(ctx, &inventory.CreateInventoryItemRequest{InventoryItem: grpcItem})
	if err != nil {
		return -1, err
	}
	item.ID = res.Id

	return res.Id, nil
}

func (a *Adapter) GetStockItem(ctx context.Context, id int64) (*domain.InventoryItem, error) {
	res, err := a.inventory.GetItem(ctx, &inventory.GetInventoryItemRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &domain.InventoryItem{
		ID:          id,
		Price:       res.InventoryItem.Price,
		StockCount:  res.InventoryItem.StockCount,
		Title:       res.InventoryItem.Title,
		Description: res.InventoryItem.Description,
		Company:     res.InventoryItem.Company,
	}, nil
}
