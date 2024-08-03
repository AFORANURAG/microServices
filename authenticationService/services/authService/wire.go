//go:build wireinject
// +build wireinject

package authService

import (
	emailServiceClient "github.com/AFORANURAG/microServices/authenticationService/clientProviders/emailServiceClient"
	queueService "github.com/AFORANURAG/microServices/authenticationService/services/queueService"
	userService "github.com/AFORANURAG/microServices/authenticationService/services/userService"
	queueServiceType "github.com/AFORANURAG/microServices/authenticationService/types/queueServiceTypes"
	userServiceTypes "github.com/AFORANURAG/microServices/authenticationService/types/userServiceTypes"

	"github.com/google/wire"
)


func InitializeAuthenticationService(phrase1 userServiceTypes.UserServicePhrase,phrase2 queueServiceType.QueueServicePhrase) AuthenticationServiceServer {
	wire.Build(NewAuthenticationServiceProvider, userService.InitializeUserService, emailServiceClient.NewEmailServiceServiceClientProvider,queueService.InitializeProducerService)
	
	return &AuthenticationServiceImpl{}
}
