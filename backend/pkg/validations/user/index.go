package validations

type (
	User struct {
		Name string ` validate:"required,min=5,max=20"`
		Age  int    ` validate:"required,gte=0"`
	}
)
