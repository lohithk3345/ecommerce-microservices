package categoryRepository

import (
	"context"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

const categoriesCollection string = "categories"

type CategoryRepository struct {
	store CategoryStore[*types.Category]
}

func NewCategoryRepository(database *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		store: NewCategoryDatabase[*types.Category](database, categoriesCollection),
	}
}

func (p CategoryRepository) InsertCategory(category *types.Category) (types.CategoryID, error) {
	result, err := p.store.insertOne(category)
	if err != nil {
		log.Println("Error", err)
		return types.EmptyString, err
	}
	log.Println("Result", result)
	return result, nil
}

func (p CategoryRepository) GetCategoryById(id types.CategoryID) (*types.Category, error) {
	result, err := p.store.findOne(H.ByID(id))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var category *types.Category
	result.Decode(&category)
	log.Println(category)
	return category, nil
}

func (p CategoryRepository) GetCategoriesByName(name string) ([]types.Category, error) {
	ctx := context.Background()
	result, err := p.store.find(ctx, H.ByName(name))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var category []types.Category
	result.All(ctx, &category)
	log.Println(&category)
	return category, nil
}

func (p CategoryRepository) GetCategoryByFilter(filter interface{}) (*types.Category, error) {
	result, err := p.store.findOne(filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var category *types.Category
	result.Decode(&category)
	log.Println(category)
	return category, nil
}
