package types

import "github.com/google/uuid"

// type ID struct {
// 	Id uuid.UUID
// }

type ID = string

// func (u *ID) SetID() {
// 	u.UUID = uuid.New().String()
// }

// func (u *ID) GetID() ID {
// 	return u.UUID
// }

// type Identifier interface {
// 	SetID()
// 	GetID() ID
// }

type Email = string

var EmptyID = uuid.Nil

const EmptyString string = string("")
