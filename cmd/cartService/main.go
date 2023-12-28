package main

import (
	cartHandler "ecommerce/api/handlers/cart"
	"ecommerce/internal/database"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func setupREST(db *mongo.Database) {
	c := cartHandler.NewCartApiHandler(db)
	router := cartHandler.SetupCartRouter(c)

	log.Printf("Starting HTTP server at: %s\n", "3003")
	http.ListenAndServe(":3003", router)
}

func main() {
	db := database.NewClient().TestDatabase()

	setupREST(db)
}
