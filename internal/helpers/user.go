package helpers

import (
	"ecommerce/types"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Helper struct {
	doc bson.M
}

var H = Helper{}

func (Helper) ByID(id types.UserID) bson.M {
	return bson.M{"_id": id}
}

func (Helper) ByEmail(email string) bson.M {
	return bson.M{"email": email}
}

func (Helper) ExtractUUIDFromInsertedID(insertedID interface{}) (uuid.UUID, error) {
	bsonBytes, err := bson.Marshal(insertedID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to marshal inserted ID to BSON: %w", err)
	}

	var insertedData types.ID
	err = bson.Unmarshal(bsonBytes, &insertedData)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to unmarshal BSON to MyData: %w", err)
	}

	return insertedData.UUID, nil
}
