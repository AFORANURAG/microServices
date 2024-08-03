package validations

type (
	User struct {
		Name  string `validate:"min=5,max=20"`
		Email string `validate:"required"`
		PhoneNumber string `validate:"required"`
		// it is you who assume that email is unique , it might not be unique ,so have checks for them and should set Unique index
		// to Email
	}
)
