package main

import (
	productHandlers "ecommerce/api/product"
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
	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		panic("err")
	}
	s := grpc.NewServer()
	p := rpcHandlers.NewProductServer(db)
	buffers.RegisterProductServiceServer(s, p)
	log.Printf("Starting gRPC server at: %s\n", "3003")

	s.Serve(lis)

	wg.Done()
}

func setupREST(db *mongo.Database) {
	p := productHandlers.NewProductApiHandler(db)
	router := productHandlers.SetupProductRouter(p)

	log.Printf("Starting HTTP server at: %s\n", "3002")
	http.ListenAndServe(":3002", router)
}

func main() {
	db := database.NewClient().TestDatabase()

	var wg sync.WaitGroup

	wg.Add(1)
	go setupGRPC(db, &wg)
	wg.Add(1)
	setupREST(db)
}
