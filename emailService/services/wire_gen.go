// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package emailService

// Injectors from wire.go:

func InitializeAuthenticationService() *EmailServiceImpl {
	emailServiceImpl := NewEmailServiceProvider()
	return emailServiceImpl
}
