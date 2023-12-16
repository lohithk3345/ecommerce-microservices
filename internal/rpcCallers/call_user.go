package rpcCallers

import (
	"context"
	buffers "ecommerce/buffers/userpb/protobuffs"
	"ecommerce/types"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserCaller struct {
	caller buffers.UserServiceClient
}

func NewUserCaller() *UserCaller {
	return &UserCaller{
		caller: connectAndGetUserClient(),
	}
}

func connectAndGetUserClient() buffers.UserServiceClient {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Could Not Connect To Server")
	}
	client := buffers.NewUserServiceClient(conn)

	return client
}

func (u UserCaller) GetUserById(id types.UserID) (*types.User, error) {
	result, err := u.caller.GetUserById(context.Background(), &buffers.GetUserByIdRequest{UserId: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user := &types.User{
		Id:      result.UserId,
		Name:    result.Name,
		Age:     int(result.Age),
		Email:   result.Email,
		Address: result.Address,
		Phone:   result.Phone,
	}

	log.Println(user)

	return user, nil

}
