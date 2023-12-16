package orderRepository

import (
	"context"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const orderCollection string = "orders"

type OrderRepository struct {
	store OrderStore[*types.Order]
}

func NewOrderRepo(database *mongo.Database) *OrderRepository {
	return &OrderRepository{
		store: NewOrderDatabase[*types.Order](database, orderCollection),
	}
}

func (u OrderRepository) InsertOrder(order *types.Order) (types.OrderID, error) {
	order.SetID()
	insertedID, err := u.store.insertOne(order)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return types.EmptyString, err
	}
	log.Println(insertedID)
	return insertedID, nil
}

func (u OrderRepository) FindOrderByID(id types.OrderID) (*types.Order, error) {
	result, err := u.store.findOne(H.ByID(id))
	if err != nil {
		log.Println("OrderRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var order *types.Order
	result.Decode(&order)
	return order, err
}

func (u OrderRepository) FindOrdersByUserID(id types.OrderID) ([]*types.Order, error) {
	ctx := context.Background()
	result, err := u.store.find(ctx, H.ByUserID(id))
	if err != nil {
		log.Println("OrderRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var orders []*types.Order
	result.All(ctx, &orders)
	log.Println(orders)
	return orders, err
}

func (u OrderRepository) FindOrderByEmail(email types.Email) (*types.Order, error) {
	log.Println("OrderRepositoryEmailLog:", H.ByEmail(email))
	result, err := u.store.findOne(H.ByEmail(email))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var order *types.Order
	result.Decode(&order)
	log.Println(order)
	return order, nil
}

func (u OrderRepository) FindByFilter(filter interface{}) (*types.Order, error) {
	result, err := u.store.findOne(filter)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var order *types.Order
	result.Decode(&order)
	return order, nil
}

func (u OrderRepository) UpdateOneById(id types.OrderID, update bson.M) {
	result, err := u.store.updateOne(H.ByID(id), bson.M{"$set": update})
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}

func (u OrderRepository) DeleteOneByID(id types.OrderID) {
	result, err := u.store.deleteOne(H.ByID(id))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}
