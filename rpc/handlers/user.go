package rpc

import (
	buffers "ecommerce/buffers/protobuffs"
	services "ecommerce/services/user"
	"ecommerce/types"
	"log"

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
	_, err := services.NewUserService(s.db).CreateUser(newUser)
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
