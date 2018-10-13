//+build wireinject

package sample

import (
	"github.com/google/go-cloud/wire"
)

var storeSet = wire.NewSet(NewOKItem, NewStore)

func initStore() (*Store, error) {
	wire.Build(storeSet)
	return nil, nil
}
