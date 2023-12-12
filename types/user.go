package types

import (
	buffers "ecommerce/buffers/protobuffs"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type UserRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    Email  `json:"email"`
	Address  string `json:"address"`
	IsActive bool   `json:"isActive"`
	Password string `json:"password"`
}

func (u UserRequest) Convert() *User {
	return &User{
		Name:     u.Name,
		Age:      u.Age,
		Address:  u.Address,
		IsActive: u.IsActive,
		Email:    u.Email,
	}
}

func ConvertRPCRequest(req *buffers.CreateUserRequest) *User {
	return &User{
		Name:    req.Name,
		Age:     int(req.Age),
		Address: req.Address,
		Email:   req.Email,
	}
}

type User struct {
	Id       UserID `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Age      int    `bson:"age"`
	Email    Email  `bson:"email"`
	Address  string `bson:"address"`
	Hash     string `bson:"hashPass"`
	IsActive bool   `bson:"isActive"`
}

func (u *User) SetID() {
	u.Id = UserID{uuid.New()}
}

func (u *User) GetID() string {
	return u.Id.UUID.String()
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) AddHash(hash string) {
	u.Hash = hash
}

type UserID struct {
	UUID uuid.UUID
}

func (u *UserID) SetID() {
	u.UUID = uuid.New()
}

func (u *UserID) GetID() uuid.UUID {
	return u.UUID
}

// func (u *User) DecodeRaw(v *mongo.SingleResult) {
// 	v.Decode(&u)
// 	log.Println("DECODE:", u.Id)
// 	uid := u.GetID()
// 	u.Id = uid
// }

func GenerateRandomEmail() string {
	rand.Seed(time.Now().UnixNano())

	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomString := make([]byte, 10)
	for i := range randomString {
		randomString[i] = allowedChars[rand.Intn(len(allowedChars))]
	}

	email := fmt.Sprintf("%s@%s", randomString, "example.domain")

	return email
}
