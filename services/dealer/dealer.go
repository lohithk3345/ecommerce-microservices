package dealerService

import (
	dealerRepository "ecommerce/internal/repositories/dealer"
	"ecommerce/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DealerServices struct {
	dealerRepository *dealerRepository.DealerRepository
}

func NewDealerService(database *mongo.Database) *DealerServices {
	return &DealerServices{
		dealerRepository: dealerRepository.NewDealerRepo(database),
	}
}

func (u DealerServices) CreateDealer(newDealer *types.Dealer) (types.ID, error) {
	return u.dealerRepository.InsertDealer(newDealer)
}

func (u DealerServices) FindById(id types.DealerID) {
	u.dealerRepository.FindDealerByID(id)
}

func (u DealerServices) FindByEmail(email types.Email) {
	u.dealerRepository.FindDealerByEmail(email)
}

func (u DealerServices) FindDealerByFilter(email types.Email) (*types.Dealer, error) {
	return u.dealerRepository.FindByFilter(bson.M{"email": email})
}
