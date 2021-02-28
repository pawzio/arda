package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pawzio/arda/service/rivendell/cmd/banner"
	"github.com/pawzio/arda/service/rivendell/cmd/serverd/rpc"
	"github.com/pawzio/arda/service/rivendell/cmd/serverd/rpc/health"
	"github.com/pawzio/arda/service/rivendell/genproto"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	banner.Print()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()

	healthpb.RegisterHealthServer(srv, health.NewServer())
	genproto.RegisterRivendellServer(srv, rpc.NewServer())

	log.Printf("starting server on port %s\n", port)
	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
