package userServices

import (
	"ecommerce/constants"
	userRepository "ecommerce/internal/repositories/user"
	"ecommerce/types"
	"slices"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServices struct {
	userRepository *userRepository.UserRepository
}

func NewUserService(database *mongo.Database) *UserServices {
	return &UserServices{
		userRepository: userRepository.NewUserRepo(database),
	}
}

func (u UserServices) CreateUser(user *types.User) (types.ID, error) {
	if user.Role == types.EmptyString || !slices.Contains(constants.Roles, user.Role) {
		user.Role = constants.Customer
	}
	return u.userRepository.InsertUser(user)
}

func (u UserServices) FindById(id types.UserID) (*types.User, error) {
	return u.userRepository.FindUserByID(id)
}

func (u UserServices) GetRoleById(id types.UserID) (types.Role, error) {
	user, err := u.userRepository.FindUserByID(id)
	if err != nil {
		return types.EmptyString, err
	}
	return user.Role, nil
}

func (u UserServices) FindByEmail(email types.Email) {
	u.userRepository.FindUserByEmail(email)
}

func (u UserServices) FindUserByFilter(email types.Email) (*types.User, error) {
	return u.userRepository.FindByFilter(bson.M{"email": email})
}
