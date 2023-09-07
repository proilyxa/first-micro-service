package repositories

import (
	"context"
	"education-project/database"
	"education-project/internal/helpers"
	"education-project/internal/models"
	"time"
)

type AuthTokenRepositoryImpl struct {
	*database.Storage
}

func NewAuthTokenRepository(storage *database.Storage) AuthTokenRepository {
	return &AuthTokenRepositoryImpl{storage}
}

func (a AuthTokenRepositoryImpl) Save(ctx context.Context, token *models.AuthToken) (*models.AuthToken, error) {
	tx, err := a.Storage.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	sql := "INSERT INTO auth_tokens( user_id, token, updated_at, created_at) values ($1, $2, $3, $4)"
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	result, err := tx.ExecContext(
		ctx,
		sql,
		token.UserId,
		token.Token,
		timestamp,
		timestamp,
	)

	helpers.PanicIfError(err)
	token.Id, err = result.LastInsertId()
	helpers.PanicIfError(err)
	token.UpdatedAt = timestamp
	token.CreatedAt = timestamp

	return token, nil
}
