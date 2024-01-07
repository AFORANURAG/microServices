package main

import (
	service "backend/apiGateway/apiGateway/pkg/services"
)

func main() {
	authService := service.InitializeAuthenticationService()
	authService.Signup()
}
