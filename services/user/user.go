package services

import (
	"ecommerce/internal/auth"
	"ecommerce/internal/repositories"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServices struct {
	userRepository *repositories.UserRepository
}

func NewUserService(database *mongo.Database) *UserServices {
	return &UserServices{
		userRepository: repositories.NewUserRepo(database),
	}
}

func (u UserServices) CreateUser(newUser *types.User) (interface{}, error) {
	hash, errPass := auth.HashPassword([]byte("Hey"))
	if errPass != nil {
		return nil, errPass
	}
	newUser.AddHash(string(hash))
	result, err := u.userRepository.InsertUser(newUser)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u UserServices) FindUser(id types.UserID) {
	u.userRepository.FindUserByID(id)
}

func (u UserServices) FindEmail(email types.Email) {
	u.userRepository.FindUserByEmail(email)
}
