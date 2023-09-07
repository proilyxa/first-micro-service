package requests

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,max=50"`
}
