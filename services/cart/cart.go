package cartServices

import (
	cartRepository "ecommerce/internal/repositories/cart"
	"ecommerce/internal/rpcCallers"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type CartServices struct {
	cartRepository *cartRepository.CartRepository
	userCaller     *rpcCallers.UserCaller
	productCaller  *rpcCallers.ProductCaller
}

func NewCartService(database *mongo.Database) *CartServices {
	return &CartServices{
		cartRepository: cartRepository.NewCartRepo(database),
		userCaller:     rpcCallers.NewUserCaller(),
		productCaller:  rpcCallers.NewProductCaller(),
	}
}

func (u *CartServices) AddToCart(userId types.UserID, productId types.ProductID) (types.CartID, error) {
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

	cart := &types.Cart{
		UserId:  user.Id,
		Product: *product,
	}

	return u.cartRepository.InsertCart(cart)
}

func (u *CartServices) FindById(id types.CartID) (*types.Cart, error) {
	return u.cartRepository.FindCartByID(id)
}

func (u *CartServices) FindByUserId(id types.CartID) ([]*types.Cart, error) {
	return u.cartRepository.FindCartsByUserID(id)
}

func (u *CartServices) RemoveItemByCartId(id types.CartID) error {
	return u.cartRepository.DeleteOneByID(id)
}

// func (u CartServices) FindCartByFilter(email types.Email) (*types.Cart, error) {
// 	return u.cartRepository.FindByFilter(bson.M{"email": email})
// }
