package registerAuth

type InputRegister struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	UserType  string `json:"userType" validate:"required"`
}
