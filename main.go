package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/geshtng/go-base-backend/db"
	"github.com/geshtng/go-base-backend/internal/handlers"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	handlers.InitAllHandlers(server)

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
