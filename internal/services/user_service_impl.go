package services

import (
	"context"
	"crypto/sha256"
	"education-project/internal/helpers"
	requests "education-project/internal/http/requests/user_requests"
	"education-project/internal/models"
	"education-project/internal/repositories"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
	AuthTokenRepo  repositories.AuthTokenRepository
}

func NewUserServiceImpl(
	userRepository repositories.UserRepository,
	authTokenRepo repositories.AuthTokenRepository,
) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		AuthTokenRepo:  authTokenRepo,
	}
}

func (u UserServiceImpl) Store(ctx context.Context, request requests.RegisterRequest) (*models.User, error) {
	hashPassword, _ := helpers.Hash(request.Password)
	user := &models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  hashPassword,
	}

	user, err := u.UserRepository.Save(ctx, user)

	return user, err
}

func (u UserServiceImpl) Login(ctx context.Context, request requests.LoginRequest) (string, error) {
	user, err := u.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if helpers.CheckHash(request.Password, user.Password) {
		return u.GenerateAuthToken(ctx, user), nil
	}

	return "", errors.New("password not valid")
}

func (u UserServiceImpl) GenerateAuthToken(ctx context.Context, user *models.User) string {
	uuidToken, err := uuid.NewRandom()
	helpers.PanicIfError(err)
	tokenString := uuidToken.String()

	h := sha256.New()
	h.Write([]byte(tokenString))
	authToken := &models.AuthToken{
		Token:  fmt.Sprintf("%x", h.Sum(nil)),
		UserId: user.Id,
	}

	authToken, err = u.AuthTokenRepo.Save(ctx, authToken)

	return tokenString
}
