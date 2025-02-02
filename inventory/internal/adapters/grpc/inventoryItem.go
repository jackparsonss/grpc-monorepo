package grpc

import (
	"context"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/application/domain"
	"github.com/jackparsonss/grpc-monorepo/proto/golang/inventory"
)

func (a *Adapter) CreateItem(ctx context.Context, request *inventory.CreateInventoryItemRequest) (*inventory.CreateInventoryItemResponse, error) {
	ie := request.InventoryItem
	newInventoryItem := domain.NewInventoryItem(ie.Price, ie.Title, ie.Company, ie.Description, ie.StockCount)

	res, err := a.api.AddInventoryItem(newInventoryItem)
	if err != nil {
		return nil, err
	}

	return &inventory.CreateInventoryItemResponse{Id: res.ID}, nil
}

func (a *Adapter) GetItem(ctx context.Context, request *inventory.GetInventoryItemRequest) (*inventory.GetInventoryItemResponse, error) {
	ie, err := a.api.GetInventoryItem(request.Id)
	if err != nil {
		return nil, err
	}

	return &inventory.GetInventoryItemResponse{
		InventoryItem: &inventory.InventoryItem{
			OptionalId:  &inventory.InventoryItem_Id{Id: ie.ID},
			Title:       ie.Title,
			Description: ie.Description,
			Company:     ie.Company,
			StockCount:  ie.StockCount,
			Price:       ie.Price,
		},
	}, nil
}
