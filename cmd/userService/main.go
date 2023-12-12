package main

import (
	apiHandlers "ecommerce/api/user"
	"ecommerce/internal/database"
	rpc "ecommerce/rpc/handlers"
	userpb "ecommerce/userpb/protobuffs"
	"log"
	"net"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func setupGRPC(db *mongo.Database, wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic("err")
	}

	s := grpc.NewServer()
	server := rpc.NewServer(db)
	userpb.RegisterUserServiceServer(s, server)
	log.Printf("Starting gRPC at: %s\n", "3000")
	s.Serve(lis)

	wg.Done()
}

func setupREST(db *mongo.Database, wg *sync.WaitGroup) {
	u := apiHandlers.NewUserApiHandler(db)
	router := apiHandlers.SetupUserRouter(u)

	log.Printf("Starting HTTP server at: %s\n", "3000")
	http.ListenAndServe(":3001", router)
	wg.Done()

}

func main() {
	db := database.NewClient().AppDatabase()

	var wg sync.WaitGroup

	wg.Add(1)
	go setupGRPC(db, &wg)
	wg.Add(1)
	go setupREST(db, &wg)

	wg.Wait()
}
