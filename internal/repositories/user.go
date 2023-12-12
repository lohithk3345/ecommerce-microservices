package repositories

import (
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection string = "users"

type UserRepository struct {
	store Store[*types.User]
}

func NewUserRepo(database *mongo.Database) *UserRepository {
	return &UserRepository{
		store: NewDatabase[*types.User](database, userCollection),
	}
}

func (u UserRepository) InsertUser(user *types.User) (*types.UserID, error) {
	user.SetID()
	insertedID, err := u.store.insertOne(user)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	return &types.UserID{UUID: insertedID.GetID()}, nil
}

func (u UserRepository) FindUserByID(id types.UserID) (*types.User, error) {
	log.Println("UserRepositoryLOG", H.ByID(id))
	result, err := u.store.findOne(H.ByID(id))
	if err != nil {
		log.Println("UserRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var user *types.User = &types.User{}
	result.Decode(&user)
	return user, err
}

func (u UserRepository) FindUserByEmail(email types.Email) {
	log.Println("UserRepositoryEmailLog:", H.ByEmail(email))
	result, err := u.store.findOne(H.ByEmail(email))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	var user types.User
	result.Decode(&user)
	log.Println(user)
}

func (u UserRepository) UpdateOneById(id types.UserID) {
	result, err := u.store.updateOne(H.ByID(id), bson.M{"$set": bson.M{"phone": 6465464664568}})
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}

func (u UserRepository) DeleteOneByID(id types.UserID) {
	result, err := u.store.deleteOne(H.ByID(id))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}

// func (u *UserRepository) UpsertOneById(id types.ID) {
// 	result, err := u.store.upsertOne(h.ByID(id), bson.M{"$set": bson.M{"phone": 6465464664568, "_id": "e2c8fd8b-39df-46e3-9a1a-81abbbb1bb46"}})
// 	if err != nil {
// 		log.Println(err.(StoreError).message)
// 		return
// 	}
// 	log.Println(result)
// }
