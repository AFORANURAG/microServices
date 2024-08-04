package validations

type (
	User struct {
		Name        string `json:"name" validate:"omitempty,min=5,max=20"`
		Email       string `json:"email" validate:"omitempty,email"`
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		// Note: Ensure that the email field is unique at the database level by setting a unique index on the Email column
	}
)
