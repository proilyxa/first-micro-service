package services

import (
	"context"
	requests "education-project/internal/http/requests/user_requests"
	"education-project/internal/models"
)

type UserService interface {
	Store(ctx context.Context, request requests.RegisterRequest) (*models.User, error)
	GenerateAuthToken(ctx context.Context, user *models.User) string
	Login(ctx context.Context, request requests.LoginRequest) (string, error)
	//Update(ctx context.Context, request request.BookUpdateRequest)
	//FindById(ctx context.Context, bookId int) response.BookResponse
	//FindAll(ctx context.Context) []response.BookResponse
}
