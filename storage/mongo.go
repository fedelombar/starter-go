// EXAMPLE
package storage

import "github.com/fedelombar/starter-go/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}
}
