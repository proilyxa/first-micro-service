package requests

type RegisterRequest struct {
	FirstName       string `json:"firstName" validate:"required,min=2,max=50"`
	LastName        string `json:"lastName" validate:"required,min=2,max=50"`
	Email           string `json:"email" validate:"required,email,max=50"`
	Password        string `json:"password" validate:"required,max=50"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
}
