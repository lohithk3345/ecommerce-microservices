package main

import (
	productHandlers "ecommerce/api/handlers/product"
	buffers "ecommerce/buffers/productpb/protobuffs"
	"ecommerce/internal/database"
	"ecommerce/internal/rpcHandlers"
	"log"
	"net"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func setupGRPC(db *mongo.Database, wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("err")
	}
	s := grpc.NewServer()
	p := rpcHandlers.NewProductServer(db)
	buffers.RegisterProductServiceServer(s, p)
	log.Printf("Starting gRPC at: %s\n", "50052")

	s.Serve(lis)

	wg.Done()
}

func setupREST(db *mongo.Database) {
	p := productHandlers.NewProductApiHandler(db)
	router := productHandlers.SetupProductRouter(p)

	log.Printf("Starting HTTP server at: %s\n", "3001")
	http.ListenAndServe(":3001", router)
}

func main() {
	db := database.NewClient().TestDatabase()

	var wg sync.WaitGroup

	wg.Add(1)
	go setupGRPC(db, &wg)
	wg.Add(1)
	setupREST(db)
}
