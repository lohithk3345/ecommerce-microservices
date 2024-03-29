package repositories

import (
	"context"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var h = helpers.Helper{}

type Database[T interface{}] struct {
	collection *mongo.Collection
}

func NewDatabase[T interface{}](database *mongo.Database, collection string) *Database[T] {
	return &Database[T]{
		collection: database.Collection(collection),
	}
}

func (d *Database[T]) insertOne(object T) (types.ID, error) {
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

func (d *Database[T]) findOne(filter interface{}) (*mongo.SingleResult, error) {
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

func (d *Database[T]) updateOne(filter interface{}, update interface{}) (reporesult.InsertResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error Store:", err.Error())
		return nil, err
	}
	if !reporesult.IsMatched(result.MatchedCount) {
		return nil, reporesult.StoreError{Code: 404, Message: "No Match Found"}
	}
	if !reporesult.IsModified(result.ModifiedCount) {
		return nil, reporesult.StoreError{Code: 409, Message: "Found match but has the value already"}
	}
	log.Println(result.MatchedCount, result.UpsertedCount, result.UpsertedID, result.ModifiedCount)
	return result.UpsertedID, nil
}

func (d *Database[T]) deleteOne(filter interface{}) (reporesult.InsertResult, error) {
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

func (d *Database[T]) find(filter interface{}) (*mongo.Cursor, error) {
	return d.collection.Find(context.Background(), filter)
}

//	func (s *Store[T]) upsertOne(filter interface{}, update interface{}) (InsertResult, error) {
//		ioptions := options.Update().SetUpsert(true)
//		result, err := s.collection.UpdateOne(context.Background(), bson.M{}, update, ioptions)
//		if err != nil {
//			log.Println("Error Store:", err.Error())
//			return nil, err
//		}
//		if err != nil {
//			log.Println("Error Store:", err.Error())
//			return nil, err
//		}
//		if !isMatched(result.MatchedCount) {
//			return nil, StoreError{Code: 404, Message: "No Match Found"}
//		}
//		if !isUpserted(result.ModifiedCount) {
//			return nil, StoreError{Code: 409, Message: "Found match but has the value already"}
//		}
//		log.Println(result.MatchedCount, result.UpsertedCount, result.UpsertedID, result.ModifiedCount)
//		return result.UpsertedID, nil
//	}
func convertDToUUID(d primitive.D) (*uuid.UUID, error) {
	// Extract the string representation of the UUID from the interface
	uuidStr, err := bson.Marshal(d)

	// Parse the string into a UUID object
	uuid, err := uuid.FromBytes(uuidStr)
	if err != nil {
		return nil, err
	}

	return &uuid, nil
}
