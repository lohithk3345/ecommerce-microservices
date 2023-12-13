package helpers

import (
	"ecommerce/types"
	"fmt"
	"log"

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

func (Helper) ExtractUUIDFromInsertedID(insertedID interface{}) (*types.ID, error) {
	bsonBytes, err := bson.Marshal(insertedID)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal inserted ID to BSON: %w", err)
	}

	var insertedData types.ID
	err = bson.Unmarshal(bsonBytes, &insertedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal BSON to MyData: %w", err)
	}

	return &insertedData, nil
}

func (h Helper) GetToken(ctx *gin.Context) (types.Token, error) {
	token := ctx.GetHeader("Authorization")
	log.Println(token)
	bearerFix := "Bearer "
	if token == "" || len(token) < len(bearerFix) || token[:len(bearerFix)] != bearerFix {
		return "", gin.Error{}
	}
	actualToken := token[len(bearerFix):]
	return actualToken, nil
}
