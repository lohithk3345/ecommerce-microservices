package orderServices

import (
	buffers "ecommerce/buffers/productpb/protobuffs"
	"ecommerce/constants"
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

	err := u.productCaller.StockUpdate(productId, buffers.StockUpdate_DEC, 1)
	if err != nil {
		return types.EmptyString, err
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

func (u OrderServices) CancelOneById(id types.OrderID, productId types.ProductID) error {
	errP := u.productCaller.StockUpdate(productId, buffers.StockUpdate_INC, 1)
	if errP != nil {
		return errP
	}

	return u.orderRepository.UpdateOneById(id, bson.M{"status": constants.OrderCancelled})
}
