package types

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type ID = string

type Email = string

type Token = string

type Update = bson.M

var EmptyID = uuid.Nil

type Role = string

const EmptyString string = string("")
