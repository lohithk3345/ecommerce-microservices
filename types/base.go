package types

import "github.com/google/uuid"

// type ID struct {
// 	Id uuid.UUID
// }

type ID struct {
	UUID uuid.UUID
}

func (u *ID) SetID() {
	u.UUID = uuid.New()
}

func (u *ID) GetID() uuid.UUID {
	return u.UUID
}

type Identifier interface {
	SetID()
	GetID() uuid.UUID
}

type Email = string

var EmptyID = uuid.Nil

const EmptyString string = string("")
