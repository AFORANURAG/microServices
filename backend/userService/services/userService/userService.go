package userService

import (
	context "context"

	"github.com/AFORANURAG/microServices/backend/userService/userRepository"
)

type UserServiceImpl struct {
	UserRepo userRepository.IUserRepository
}

func (u *UserServiceImpl) GetUserByName(context.Context, *Request) (*Response, error) {
	user, err := u.UserRepo.GetUserByName("Anurag")
	return &Response{Name: user.Name}, err

}

func (u *UserServiceImpl) GetUserById(context.Context, *Request) (*Response, error) {
	user, err := u.UserRepo.GetUserByRowId("Anurag")
	return &Response{Name: user.Name}, err
}

func (u *UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}
func NewUserServiceProvider(urepo *userRepository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: urepo}
}
