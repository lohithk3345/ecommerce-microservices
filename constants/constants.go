package constants

import "ecommerce/types"

const API_KEY string = "x-api-key"
const TOKEN_SECRET string = "TOKEN_SECRET"

const (
	Customer types.Role = "Customer"
	Dealer   types.Role = "Dealer"
)

var Roles []types.Role = []types.Role{Customer, Dealer}
