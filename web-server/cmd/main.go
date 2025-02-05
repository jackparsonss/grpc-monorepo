package main

import (
	"github.com/jackparsonss/grpc-monorepo/web-server/config"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/adapters/inventory"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/adapters/web"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/api"
	"log"
)

func main() {
	inventoryAdapter, err := inventory.NewAdapter(config.GetInventoryServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(inventoryAdapter)
	apiAdapter := web.NewAdapter(application, config.GetApplicationPort())
	apiAdapter.Run()
}
