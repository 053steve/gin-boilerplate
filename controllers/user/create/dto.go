package createUser

type CreateUserDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserType  string `json:"userType" validate:"required"`
	Active    bool   `json:"active;"`
}
