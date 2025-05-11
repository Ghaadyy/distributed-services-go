package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Ghaadyy/proglog/internal/server"
)

func main() {
	// srv := server.NewHTTPServer(":8080")
	// log.Fatal(srv.ListenAndServe())

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv, err := server.NewGRPCServer(&server.Config{
		// CommitLog: server.CommitLog,
	})
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
	fmt.Printf("started listening on port %d\r\n", 8080)
	srv.Serve(lis)
}
