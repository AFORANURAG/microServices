//go:build wireinject
// +build wireinject

package queueservice

import (
	"github.com/google/wire"
)

func InitializeProducerService(phrase string) *Producer {
	wire.Build(NewProducerServiceProvider)
	return &Producer{}
}
