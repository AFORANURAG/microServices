package service

import "fmt"

type IEmailService interface {
	Signup() string
}

type EmailServiceImpl struct {
	authServiceUri string
}

func (a *EmailServiceImpl) Signup() string {
	fmt.Println("Fuck off")
	return "hello"
}

func EmailServiceProvider(u string) *EmailServiceImpl {
	return &EmailServiceImpl{u}
}
