package authenticationService

type IAuthenticationService interface {
	Signup() string
}

type AuthenticationServiceImpl struct {
}

func (a *AuthenticationServiceImpl) Signup() string {
	return "hello world , I am in Sign up"
}

func AuthenticationServiceProvider() IAuthenticationService {
	return &AuthenticationServiceImpl{}
}
