package configs

import (
	"backend/pkg/services"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/dig"
)

var Container = dig.New()

func InjectServices() {
	// Container.Provide(authenticationService.AuthenticationServiceProvider)
	if err := Container.Provide(services.RegisterServices); err != nil {
		fmt.Println("error while providing the services")
		log.Info(err)
	}
	// if err := Container.Invoke(func(services *services.Services) {
	// 	fmt.Println(services.AuthenticationService.Signup())
	// }); err != nil {
	// 	log.Info(err)
	// }
}
