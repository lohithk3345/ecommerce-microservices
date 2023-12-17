package rpcHandlers

import (
	buffers "ecommerce/buffers/userpb/protobuffs"
	"ecommerce/constants"
	"ecommerce/internal/auth"
	userServices "ecommerce/services/user"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	db *mongo.Database
	buffers.UnimplementedUserServiceServer
}

func NewUserServer(db *mongo.Database) *UserServer {
	return &UserServer{db: db}
}

func (s *UserServer) CreateUser(ctx context.Context, req *buffers.CreateUserRequest) (*buffers.CreateUserResponse, error) {
	newUser := types.ConvertUserRPCRequest(req)
	hash, errPass := auth.HashPassword([]byte(req.Password))
	if errPass != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	newUser.AddHash(string(hash))
	newUser.IsActive = true
	_, err := userServices.NewUserService(s.db).CreateUser(newUser)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "The Email Already Exists")
	}
	logger := log.Default()
	logger.Println("Success")

	return &buffers.CreateUserResponse{
		Status: 200,
		Body:   "User Created",
	}, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *buffers.GetUserByIdRequest) (*buffers.GetUserByIdResponse, error) {
	user, err := userServices.NewUserService(s.db).FindById(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User "+codes.NotFound.String())
	}
	if user.Role != constants.Customer {
		return nil, status.Error(codes.NotFound, "User "+codes.NotFound.String())
	}

	log.Println("Found User")

	return &buffers.GetUserByIdResponse{
		UserId:  user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Age:     int32(user.Age),
		Address: user.Address,
		Phone:   user.Phone,
	}, nil
}
