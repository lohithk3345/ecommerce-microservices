package userRepository

import (
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection string = "users"

type UserRepository struct {
	store UserStore[*types.User]
}

func NewUserRepo(database *mongo.Database) *UserRepository {
	return &UserRepository{
		store: NewUserDatabase[*types.User](database, userCollection),
	}
}

func (u UserRepository) InsertUser(user *types.User) (types.ID, error) {
	user.SetID()
	insertedID, err := u.store.insertOne(user)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return types.EmptyString, err
	}
	log.Println(insertedID)
	return insertedID, nil
}

func (u UserRepository) FindUserByID(id types.UserID) (*types.User, error) {
	result, err := u.store.findOne(H.ByID(id))
	if err != nil {
		log.Println("UserRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var user *types.User
	result.Decode(&user)
	return user, err
}

func (u UserRepository) FindUserByEmail(email types.Email) (*types.User, error) {
	log.Println("UserRepositoryEmailLog:", H.ByEmail(email))
	result, err := u.store.findOne(H.ByEmail(email))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var user *types.User
	result.Decode(&user)
	log.Println(user)
	return user, nil
}

func (u UserRepository) FindByFilter(filter interface{}) (*types.User, error) {
	result, err := u.store.findOne(filter)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var user *types.User
	result.Decode(&user)
	return user, nil
}

func (u UserRepository) UpdateOneById(id types.UserID, update bson.M) {
	result, err := u.store.updateOne(H.ByID(id), bson.M{"$set": update})
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
