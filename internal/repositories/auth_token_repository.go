package repositories

import (
	"context"
	"education-project/internal/models"
)

type AuthTokenRepository interface {
	Save(ctx context.Context, token *models.AuthToken) (*models.AuthToken, error)
}
