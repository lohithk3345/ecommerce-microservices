package rpcCallers

import (
	"context"
	buffers "ecommerce/buffers/productpb/protobuffs"
	"ecommerce/types"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductCaller struct {
	caller buffers.ProductServiceClient
}

func NewProductCaller() *ProductCaller {
	return &ProductCaller{
		caller: connectAndGetProductClient(),
	}
}

func connectAndGetProductClient() buffers.ProductServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Could Not Connect To Server")
	}
	client := buffers.NewProductServiceClient(conn)

	return client
}

func (p ProductCaller) GetProductById(id types.ProductID) (*types.Product, error) {
	result, err := p.caller.GetProductById(context.Background(), &buffers.GetProductByIdRequest{ProductId: id})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	product := &types.Product{
		Id:          result.Id,
		Name:        result.Name,
		Price:       result.Price,
		Description: result.Description,
		DealerId:    result.DealerId,
	}
	log.Println("Product", product)

	return product, err
}

func (p ProductCaller) StockUpdate(id types.ProductID, operaion buffers.StockUpdate, number int32) error {
	req := &buffers.UpdateStockByIdRequest{
		ProductId: id,
		Operation: operaion,
		ByNumber:  uint32(number),
	}
	result, err := p.caller.StockUpdateById(context.Background(), req)
	if err != nil {
		log.Println(err)
		return err
	}

	if result.Status == buffers.Status_FAILED {
		return errors.New("Operation Failed")
	}

	return nil
}
