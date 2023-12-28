package cartRepository

import (
	"context"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const cartCollection string = "carts"

type CartRepository struct {
	store CartStore[*types.Cart]
}

func NewCartRepo(database *mongo.Database) *CartRepository {
	return &CartRepository{
		store: NewCartDatabase[*types.Cart](database, cartCollection),
	}
}

func (u CartRepository) InsertCart(cart *types.Cart) (types.CartID, error) {
	cart.SetID()
	cart.SetCreatedTimestamp()
	// cart.SetUpdatedTimestamp()
	insertedID, err := u.store.insertOne(cart)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return types.EmptyString, err
	}
	log.Println(insertedID)
	return insertedID, nil
}

func (u CartRepository) FindCartByID(id types.CartID) (*types.Cart, error) {
	result, err := u.store.findOne(H.ByID(id))
	if err != nil {
		log.Println("CartRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var cart *types.Cart
	result.Decode(&cart)
	log.Println(cart)
	return cart, err
}

func (u CartRepository) FindCartsByUserID(id types.CartID) ([]*types.Cart, error) {
	ctx := context.Background()
	result, err := u.store.find(ctx, H.ByUserID(id))
	if err != nil {
		log.Println("CartRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var carts []*types.Cart
	result.All(ctx, &carts)
	log.Println(carts)
	return carts, err
}

func (u CartRepository) FindCartByEmail(email types.Email) (*types.Cart, error) {
	log.Println("CartRepositoryEmailLog:", H.ByEmail(email))
	result, err := u.store.findOne(H.ByEmail(email))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var cart *types.Cart
	result.Decode(&cart)
	log.Println(cart)
	return cart, nil
}

func (u CartRepository) FindByFilter(filter interface{}) (*types.Cart, error) {
	result, err := u.store.findOne(filter)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var cart *types.Cart
	result.Decode(&cart)
	return cart, nil
}

func (u CartRepository) UpdateOneById(id types.CartID, update bson.M) error {
	err := u.store.updateOne(H.ByID(id), bson.M{"$set": update})
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return err
	}
	return nil
}

func (u CartRepository) DeleteOneByID(id types.CartID) error {
	result, err := u.store.deleteOne(H.ByID(id))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return err
	}
	log.Println(result)
	return nil
}
