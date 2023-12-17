package categoriesServices

import (
	categoryRepository "ecommerce/internal/repositories/product/category"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryServices struct {
	repository categoryRepository.CategoryRepository
}

func NewCategoryService(database *mongo.Database) *CategoryServices {
	return &CategoryServices{
		repository: *categoryRepository.NewCategoryRepository(database),
	}
}

func (p CategoryServices) CreateCategory(product *types.Category) (types.CategoryID, error) {
	return p.repository.InsertCategory(product)
}

func (u CategoryServices) FindCategoryById(id types.CategoryID) (*types.Category, error) {
	return u.repository.GetCategoryById(id)
}

func (p CategoryServices) FindCategorysByName(name string) ([]types.Category, error) {
	return p.repository.GetCategoriesByName(name)
}
