package validations

type (
	User struct {
		Name  string ` validate:"required,min=5,max=20"`
		Email string ` validate:"required"`
	}
)
