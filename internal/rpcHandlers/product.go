package rpcHandlers

import (
	"context"
	buffers "ecommerce/buffers/productpb/protobuffs"
	productServices "ecommerce/services/product"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	db *mongo.Database
	buffers.UnimplementedProductServiceServer
}

func NewProductServer(db *mongo.Database) *ProductServer {
	return &ProductServer{db: db}
}

func (p *ProductServer) GetProductById(ctx context.Context, req *buffers.GetProductByIdRequest) (*buffers.GetProductByIdResponse, error) {
	product := types.ConvertProductRPCRequest(req)
	result, err := productServices.NewProductService(p.db).FindProductById(product.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Product Not Found")
	}
	response := &buffers.GetProductByIdResponse{Id: result.Id, Name: result.Name,
		Price: result.Price, Description: result.Description, DealerId: result.DealerId}
	return response, nil
}
