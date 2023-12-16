package orderServices

import (
	orderRepository "ecommerce/internal/repositories/order"
	"ecommerce/internal/rpcCallers"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderServices struct {
	orderRepository *orderRepository.OrderRepository
	userCaller      *rpcCallers.UserCaller
	productCaller   *rpcCallers.ProductCaller
}

func NewOrderService(database *mongo.Database) *OrderServices {
	return &OrderServices{
		orderRepository: orderRepository.NewOrderRepo(database),
		userCaller:      rpcCallers.NewUserCaller(),
		productCaller:   rpcCallers.NewProductCaller(),
	}
}

func (u OrderServices) CreateOrder(userId types.UserID, productId types.ProductID) (types.OrderID, error) {
	user, errUser := u.userCaller.GetUserById(userId)
	if errUser != nil {
		log.Println(errUser)
		return types.EmptyString, errUser
	}

	product, errProduct := u.productCaller.GetProductById(productId)
	if errProduct != nil {
		log.Println(errProduct)
		return types.EmptyString, errProduct
	}
	log.Println("Product", product)

	order := &types.Order{
		UserId:    user.Id,
		Address:   user.Address,
		Name:      user.Name,
		Price:     product.Price,
		DealerId:  product.DealerId,
		ProductId: product.Id,
		Email:     user.Email,
	}

	return u.orderRepository.InsertOrder(order)
}

func (u OrderServices) FindById(id types.OrderID) (*types.Order, error) {
	return u.orderRepository.FindOrderByID(id)
}

func (u OrderServices) FindByUserId(id types.OrderID) ([]*types.Order, error) {
	return u.orderRepository.FindOrdersByUserID(id)
}

func (u OrderServices) FindOrderByFilter(email types.Email) (*types.Order, error) {
	return u.orderRepository.FindByFilter(bson.M{"email": email})
}
