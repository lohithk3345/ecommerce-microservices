package repositories

import (
	"ecommerce/internal/helpers"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NOT_FOUND = iota + 1
	DUPLICATE
	CONN_ERR
)

var H = helpers.H

type Store[T interface{}] interface {
	insertOne(object T) (*types.ID, error)
	findOne(filter interface{}) (*mongo.SingleResult, error)
	updateOne(filter interface{}, update interface{}) (reporesult.InsertResult, error)
	deleteOne(filter interface{}) (reporesult.InsertResult, error)
}
