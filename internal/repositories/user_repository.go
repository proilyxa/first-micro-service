package repositories

import (
	"context"
	"education-project/internal/models"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user models.User)
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) (models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByAuthToken(ctx context.Context, token string) (*models.User, error)
	FindAll(ctx context.Context) []models.User
}
