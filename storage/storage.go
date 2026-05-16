package storage

import "github.com/fedelombar/starter-go/types"

type Storer interface {
	Get(int) *types.User
}
