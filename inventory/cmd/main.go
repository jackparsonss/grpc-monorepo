package main

import (
	"github.com/jackparsonss/grpc-monorepo/inventory/config"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/adapters/db"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/adapters/grpc"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/application/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDatabaseURL())
	if err != nil {
		log.Fatalf("Failed to connect to database, error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	log.Printf("Starting gRPC Server on port: %d\n", config.GetApplicationPort())
	grpcAdapter.Run()
}
