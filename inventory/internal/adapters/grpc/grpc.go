package grpc

import (
	"fmt"
	"github.com/jackparsonss/grpc-monorepo/inventory/config"
	"github.com/jackparsonss/grpc-monorepo/inventory/internal/ports"
	"github.com/jackparsonss/grpc-monorepo/proto/golang/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api  ports.APIPort
	port int
	inventory.UnsafeInventoryServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a *Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to list on port %d: error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	inventory.RegisterInventoryServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve on port %d: error: %v", a.port, err)
	}
}
