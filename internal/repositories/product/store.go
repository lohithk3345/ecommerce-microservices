package productRepository

import (
	"context"
	"ecommerce/internal/helpers"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NOT_FOUND = iota + 1
	DUPLICATE
	CONN_ERR
)

var H = helpers.H

type ProductStore[T interface{}] interface {
	insertOne(object T) (types.ID, error)
	findOne(filter interface{}) (*mongo.SingleResult, error)
	updateOne(filter interface{}, update interface{}) error
	deleteOne(filter interface{}) (reporesult.InsertResult, error)
	find(ctx context.Context, filter interface{}) (*mongo.Cursor, error)
}

type ProductDatabase[T interface{}] struct {
	collection *mongo.Collection
}

func NewProductDatabase[T interface{}](database *mongo.Database, collection string) *ProductDatabase[T] {
	return &ProductDatabase[T]{
		collection: database.Collection(collection),
	}
}

func (d *ProductDatabase[T]) insertOne(object T) (types.ProductID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := d.collection.InsertOne(ctx, object)
	if err != nil {
		errIn := err.(mongo.WriteException)
		log.Println(errIn.WriteErrors[0].Code)
		for _, writeError := range errIn.WriteErrors {
			log.Println(writeError.Code)
		}
		// log.Println(err.Error())
		return types.EmptyString, reporesult.StoreError{Code: 204, Message: err.Error()}
	}
	log.Println("ID", result.InsertedID)
	// id, _ := H.ExtractUUIDFromInsertedID(result.InsertedID)
	id := result.InsertedID.(types.ID)
	log.Println(id)
	return id, nil
}

func (d *ProductDatabase[T]) findOne(filter interface{}) (*mongo.SingleResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		log.Println("BaseError:", result.Err().Error())
		return nil, reporesult.StoreError{Code: 404, Message: "Not Found"}
	}
	if result == nil {
		log.Println("BaseResult: Result Nil")
	}
	return result, nil
}

func (d *ProductDatabase[T]) updateOne(filter interface{}, update interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error Store:", err.Error())
		return err
	}
	if !reporesult.IsMatched(result.MatchedCount) {
		return reporesult.StoreError{Code: 404, Message: "No Match Found"}
	}
	if !reporesult.IsModified(result.ModifiedCount) {
		return reporesult.StoreError{Code: 409, Message: "Found match but has the value already"}
	}
	log.Println(result.MatchedCount, result.UpsertedCount, result.UpsertedID, result.ModifiedCount)
	return nil
}

func (d *ProductDatabase[T]) deleteOne(filter interface{}) (reporesult.InsertResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err.Error())
		return nil, reporesult.StoreError{Code: 500, Message: err.Error()}
	}
	if !reporesult.IsDeletedCount(result.DeletedCount) {
		return nil, reporesult.StoreError{Code: 204, Message: "Did Not Delete As DO Does Not Exist"}
	}
	log.Println(result.DeletedCount)
	return reporesult.IsDeletedCount(result.DeletedCount), nil
}

func (d *ProductDatabase[T]) find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	return d.collection.Find(context.Background(), filter)
}
