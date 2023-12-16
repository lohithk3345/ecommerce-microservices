package productServices

import (
	productRepository "ecommerce/internal/repositories/product"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServices struct {
	repository productRepository.ProductRepository
}

func NewProductService(database *mongo.Database) *ProductServices {
	return &ProductServices{
		repository: *productRepository.NewProductRepository(database),
	}
}

func (p ProductServices) CreateProduct(product *types.Product) (types.ProductID, error) {
	return p.repository.InsertProduct(product)
}

func (u ProductServices) FindProductById(id types.ProductID) (*types.Product, error) {
	return u.repository.GetProductById(id)
}

func (p ProductServices) FindProductsByName(name string) ([]types.Product, error) {
	return p.repository.GetProductsByName(name)
}
