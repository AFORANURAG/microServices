package userService

import "fmt"

type IUserService interface {
	GetUser(id string, name string) string
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUser(id string, name string) string {
	fmt.Println("Fuck You Twice")
	return "Hello World!"
}

func UserServiceProvider() *UserServiceImpl {
	return &UserServiceImpl{}
}
