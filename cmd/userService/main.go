package main

import (
	userHandlers "ecommerce/api/handlers/user"
	buffers "ecommerce/buffers/userpb/protobuffs"
	"ecommerce/internal/database"
	rpcHandlers "ecommerce/internal/rpcHandlers"
	"log"
	"net"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func setupGRPC(db *mongo.Database, wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("err")
	}

	s := grpc.NewServer()
	server := rpcHandlers.NewUserServer(db)
	buffers.RegisterUserServiceServer(s, server)
	log.Printf("Starting gRPC at: %s\n", "50051")
	s.Serve(lis)

	wg.Done()
}

func setupREST(db *mongo.Database, wg *sync.WaitGroup) {
	u := userHandlers.NewUserApiHandler(db)
	router := userHandlers.SetupUserRouter(u)

	log.Printf("Starting HTTP server at: %s\n", "3000")
	http.ListenAndServe(":3000", router)
	wg.Done()

}

func main() {
	db := database.NewClient().TestDatabase()

	var wg sync.WaitGroup

	wg.Add(1)
	go setupGRPC(db, &wg)
	wg.Add(1)
	setupREST(db, &wg)

	wg.Wait()
}
