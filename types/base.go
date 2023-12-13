package types

import "github.com/google/uuid"

type ID = string

type Email = string

type Token = string

var EmptyID = uuid.Nil

const EmptyString string = string("")
