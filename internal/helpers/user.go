package helpers

import (
	"ecommerce/types"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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

func (Helper) ByName(name string) bson.M {
	return bson.M{"name": name}
}

func (Helper) ByUserID(id types.UserID) bson.M {
	return bson.M{"userId": id}
}

func (Helper) ByProductID(id types.UserID) bson.M {
	return bson.M{"productId": id}
}

func (Helper) SetTimestamp() time.Time {
	return time.Now().UTC()
}

func (Helper) ExtractUUIDFromInsertedID(insertedID interface{}) (*types.ID, error) {
	log.Println("START")
	bsonBytes, err := bson.Marshal(insertedID)
	if err != nil {
		log.Println("Error")
		return nil, fmt.Errorf("failed to marshal inserted ID to BSON: %w", err)
	}

	log.Println(bsonBytes)

	var insertedData types.ID
	err = bson.Unmarshal(bsonBytes, &insertedData)
	log.Println(insertedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON to MyData: %w", err)
	}

	return &insertedData, nil
}

func (h Helper) GetToken(ctx *gin.Context) (types.Token, error) {
	token := ctx.GetHeader("Authorization")
	bearerFix := "Bearer "
	if token == types.EmptyString || len(token) < len(bearerFix) || token[:len(bearerFix)] != bearerFix {
		return types.EmptyString, gin.Error{}
	}
	actualToken := token[len(bearerFix):]
	return actualToken, nil
}
