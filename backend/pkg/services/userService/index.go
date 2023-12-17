package userService

type IUserService interface {
	GetUser(id string, name string) string
}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) GetUser(id string, name string) string {
	return "Hello World!"
}

func UserServiceProvider() IUserService {
	return &UserServiceImpl{}
}
