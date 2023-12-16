package dealerRepository

import (
	"ecommerce/constants"
	reporesult "ecommerce/internal/repositories/repo_result"
	"ecommerce/types"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const dealerCollection string = "dealers"

type DealerRepository struct {
	store DealerStore[*types.Dealer]
}

func NewDealerRepo(database *mongo.Database) *DealerRepository {
	return &DealerRepository{
		store: NewDealerDatabase[*types.Dealer](database, dealerCollection),
	}
}

func (u DealerRepository) InsertDealer(dealer *types.Dealer) (types.DealerID, error) {
	dealer.SetID()
	dealer.Role = constants.Customer
	insertedID, err := u.store.insertOne(dealer)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return types.EmptyString, err
	}
	log.Println(insertedID)
	return insertedID, nil
}

func (u DealerRepository) FindDealerByID(id types.DealerID) (*types.Dealer, error) {
	result, err := u.store.findOne(H.ByID(id))
	if err != nil {
		log.Println("DealerRepository", err.(reporesult.StoreError).Message)
		return nil, err
	}

	var dealer *types.Dealer
	result.Decode(&dealer)
	return dealer, err
}

func (u DealerRepository) FindDealerByEmail(email types.Email) (*types.Dealer, error) {
	log.Println("DealerRepositoryEmailLog:", H.ByEmail(email))
	result, err := u.store.findOne(H.ByEmail(email))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var dealer *types.Dealer
	result.Decode(&dealer)
	log.Println(dealer)
	return dealer, nil
}

func (u DealerRepository) FindByFilter(filter interface{}) (*types.Dealer, error) {
	result, err := u.store.findOne(filter)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return nil, err
	}
	var dealer *types.Dealer
	result.Decode(&dealer)
	return dealer, nil
}

func (u DealerRepository) UpdateOneById(id types.DealerID, update bson.M) {
	result, err := u.store.updateOne(H.ByID(id), update)
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}

func (u DealerRepository) DeleteOneByID(id types.DealerID) {
	result, err := u.store.deleteOne(H.ByID(id))
	if err != nil {
		log.Println(err.(reporesult.StoreError).Message)
		return
	}
	log.Println(result)
}
