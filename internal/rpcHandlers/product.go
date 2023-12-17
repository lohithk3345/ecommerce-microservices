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
	if result.Stock <= 0 {
		return nil, status.Error(codes.NotFound, "Product Stock Empty")
	}
	response := &buffers.GetProductByIdResponse{Id: result.Id, Name: result.Name,
		Price: result.Price, Description: result.Description, DealerId: result.DealerId}
	return response, nil
}

func (p *ProductServer) StockUpdateById(ctx context.Context, req *buffers.UpdateStockByIdRequest) (*buffers.UpdateStockByIdResponse, error) {
	product := types.ConvertProductRPCIncRequest(req)
	operation := req.Operation
	ps := productServices.NewProductService(p.db)
	result, err := ps.FindProductById(product.Id)
	if err != nil {
		return &buffers.UpdateStockByIdResponse{
			Status: buffers.Status_FAILED,
		}, status.Error(codes.NotFound, "Product Not Found")
	}
	if result.Stock <= 0 {
		return &buffers.UpdateStockByIdResponse{
			Status: buffers.Status_FAILED,
		}, status.Error(codes.NotFound, "Product Stock Empty")
	}
	if operation == buffers.StockUpdate_INC {
		uErr := ps.IncrementStockById(product.Id, int16(req.ByNumber))
		if uErr != nil {
			return &buffers.UpdateStockByIdResponse{
				Status: buffers.Status_FAILED,
			}, status.Error(codes.Canceled, "Could Not Increment")
		}
	}

	if operation == buffers.StockUpdate_DEC {
		uErr := ps.DecrementStockById(product.Id, int16(req.ByNumber))
		if uErr != nil {
			return &buffers.UpdateStockByIdResponse{
				Status: buffers.Status_FAILED,
			}, status.Error(codes.Canceled, "Could Not Decrement")
		}
	}

	return &buffers.UpdateStockByIdResponse{
		Status: buffers.Status_SUCCESS,
	}, nil
}
