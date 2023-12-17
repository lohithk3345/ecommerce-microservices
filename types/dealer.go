package types

import "github.com/google/uuid"

type DealerID = ID

type Dealer struct {
	Id       DealerID `bson:"_id,omitempty"`
	Name     string   `bson:"name"`
	Age      int      `bson:"age"`
	Email    Email    `bson:"email"`
	Address  string   `bson:"address"`
	Hash     string   `bson:"hashPass"`
	IsActive bool     `bson:"isActive"`
	Role     string   `bson:"role"`
}

func (d *Dealer) SetID() {
	d.Id = uuid.New().String()
}

type DealerRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    Email  `json:"email"`
	Address  string `json:"address"`
	IsActive bool   `json:"isActive"`
	Password string `json:"password"`
}
