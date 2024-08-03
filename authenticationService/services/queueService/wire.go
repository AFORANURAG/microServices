//go:build wireinject
// +build wireinject

package queueservice

import (
	queueServiceType "github.com/AFORANURAG/microServices/authenticationService/types/queueServiceTypes"
	"github.com/google/wire"
)

func InitializeProducerService(phrase queueServiceType.QueueServicePhrase ) *Producer {
	wire.Build(NewProducerServiceProvider,NewQueueServiceProvider)
	return &Producer{}
}
