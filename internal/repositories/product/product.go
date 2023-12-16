package productRepository

import (
	"context"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

const productCollection string = "products"

type ProductRepository struct {
	store ProductStore[*types.Product]
}

func NewProductRepository(database *mongo.Database) *ProductRepository {
	return &ProductRepository{
		store: NewProductDatabase[*types.Product](database, productCollection),
	}
}

func (p ProductRepository) InsertProduct(product *types.Product) (types.ProductID, error) {
	product.SetID()
	result, err := p.store.insertOne(product)
	if err != nil {
		log.Println("Error", err)
		return types.EmptyString, err
	}
	log.Println("Result", result)
	return result, nil
}

func (p ProductRepository) GetProductById(id types.ProductID) (*types.Product, error) {
	result, err := p.store.findOne(H.ByID(id))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var product *types.Product
	result.Decode(&product)
	return product, nil
}

func (p ProductRepository) GetProductsByName(name string) ([]types.Product, error) {
	ctx := context.Background()
	result, err := p.store.find(ctx, H.ByName(name))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var product []types.Product
	result.All(ctx, &product)
	log.Println(&product)
	return product, nil
}

func (p ProductRepository) GetProductByFilter(filter interface{}) (*types.Product, error) {
	result, err := p.store.findOne(filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var product *types.Product
	result.Decode(&product)
	log.Println(product)
	return product, nil
}

func (p ProductRepository) GetProductsByFilter(filter interface{}) ([]types.Product, error) {
	ctx := context.Background()
	result, err := p.store.find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var product []types.Product
	result.All(ctx, &product)
	log.Println(&product)
	return product, nil
}

func (p ProductRepository) GetProductsByCategory(filter interface{}) ([]types.Product, error) {
	ctx := context.Background()
	result, err := p.store.find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var product []types.Product
	result.All(ctx, &product)
	log.Println(&product)
	return product, nil
}
