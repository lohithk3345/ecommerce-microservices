package main

import (
	orderHandler "ecommerce/api/order"
	"ecommerce/internal/database"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func setupREST(db *mongo.Database) {
	o := orderHandler.NewOrderApiHandler(db)
	router := orderHandler.SetupOrderRouter(o)

	log.Printf("Starting HTTP server at: %s\n", "3004")
	http.ListenAndServe(":3004", router)
}

func main() {
	db := database.NewClient().TestDatabase()

	setupREST(db)
}
