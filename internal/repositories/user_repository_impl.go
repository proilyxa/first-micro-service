package repositories

import (
	"context"
	"crypto/sha256"
	"education-project/database"
	"education-project/internal/helpers"
	"education-project/internal/models"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"time"
)

type UserRepositoryImpl struct {
	*database.Storage
}

func NewUserRepository(storage *database.Storage) UserRepository {
	return &UserRepositoryImpl{storage}
}

func (u UserRepositoryImpl) Save(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := u.Storage.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	sql := "INSERT INTO users(first_name, last_name, email, password, updated_at, created_at) values ($1, $2, $3, $4, $5, $6)"

	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	result, err := tx.ExecContext(
		ctx,
		sql,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		timestamp,
		timestamp,
	)

	if err != nil {
		errVal := err.(sqlite3.Error).ExtendedCode
		if errVal == sqlite3.ErrConstraintUnique {
			return nil, fmt.Errorf("This email alredy exist")
		}

		panic(err)
	}

	user.Id, err = result.LastInsertId()
	helpers.PanicIfError(err)
	user.UpdatedAt = timestamp
	user.CreatedAt = timestamp

	return user, nil
}

func (u UserRepositoryImpl) Update(ctx context.Context, user models.User) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) Delete(ctx context.Context, userId int) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) FindById(ctx context.Context, userId int) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	tx, err := u.Storage.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	sql := "select * from users where email=$1"
	result, errQuery := tx.QueryContext(ctx, sql, email)
	helpers.PanicIfError(errQuery)
	defer result.Close()

	user := &models.User{}

	if result.Next() {
		err := result.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		helpers.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (u UserRepositoryImpl) FindByAuthToken(ctx context.Context, token string) (*models.User, error) {
	tx, err := u.Storage.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	sql := `select
    		users.id,
    		users.first_name,
    		users.last_name,
    		users.email,
    		users.password,
    		auth_tokens.token,
    		users.updated_at,
    		users.created_at
			from users
			    join auth_tokens on auth_tokens.user_id = users.id
			where auth_tokens.token = $1
			limit 1`

	h := sha256.New()
	h.Write([]byte(token))

	result, errQuery := tx.QueryContext(ctx, sql, fmt.Sprintf("%x", h.Sum(nil)))
	helpers.PanicIfError(errQuery)
	defer result.Close()

	user := &models.User{}
	if result.Next() {
		err := result.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Token,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		helpers.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (u UserRepositoryImpl) FindAll(ctx context.Context) []models.User {
	//TODO implement me
	panic("implement me")
}
