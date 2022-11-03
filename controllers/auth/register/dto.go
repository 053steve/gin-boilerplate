package registerAuth

type RegisterDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserType  string `json:"userType" validate:"required"`
}
