package rpcHandlers

import (
	buffers "ecommerce/buffers/protobuffs"
	"ecommerce/internal/auth"
	services "ecommerce/services/user"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	db *mongo.Database
	buffers.UnimplementedUserServiceServer
}

func NewServer(db *mongo.Database) *Server {
	return &Server{db: db}
}

func (s *Server) CreateUser(ctx context.Context, req *buffers.CreateUserRequest) (*buffers.CreateUserResponse, error) {
	newUser := types.ConvertRPCRequest(req)
	hash, errPass := auth.HashPassword([]byte(req.Password))
	if errPass != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	newUser.AddHash(string(hash))
	newUser.IsActive = true
	_, err := services.NewUserService(s.db).CreateUser(newUser)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "The Email Already Exists")
	}

	return &buffers.CreateUserResponse{
		Status: 200,
		Body:   "User Created",
	}, nil

}
