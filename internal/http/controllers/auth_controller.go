package controllers

import (
	"education-project/internal/api/response"
	"education-project/internal/api/valid"
	userRequests "education-project/internal/http/requests/user_requests"
	"education-project/internal/pkg/logger/sl"
	"education-project/internal/services"
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
	"net/http"
)

type AuthController struct {
	UserService services.UserService
}

func NewAuthController(userService services.UserService) *AuthController {
	return &AuthController{UserService: userService}
}

func validate(resp http.ResponseWriter, req *http.Request, request interface{}) bool {
	validatorObj := valid.NewValidator()
	err := validatorObj.Struct(request)
	if err != nil {
		var validateErr validator.ValidationErrors
		errors.As(err, &validateErr)
		render.Status(req, http.StatusUnprocessableEntity)
		render.JSON(resp, req, response.ValidationError(validateErr))

		return false
	}

	return true
}

func decodeRequest(resp http.ResponseWriter, req *http.Request, op string, request interface{}) bool {
	err := render.DecodeJSON(req.Body, &request)
	if err != nil {
		sl.LoggerFromContext(req.Context()).Warn(err.Error(), slog.String("op", op))
		render.JSON(resp, req, response.Error("Failed to decode request"))
		return false
	}

	return true
}

func (controller *AuthController) Register(resp http.ResponseWriter, req *http.Request) {
	const op = "controllers.auth_controller.Register"

	var regRequest userRequests.RegisterRequest
	if decodeRequest(resp, req, op, &regRequest) == false {
		return
	}

	if validate(resp, req, &regRequest) == false {
		return
	}

	user, err := controller.UserService.Store(req.Context(), regRequest)
	if err != nil {
		render.Status(req, http.StatusUnprocessableEntity)
		render.JSON(resp, req, response.ValidationErrors(err.Error(), map[string][]string{
			"email": {err.Error()},
		}))

		return
	}

	authToken := controller.UserService.GenerateAuthToken(req.Context(), user)
	respStruct := struct {
		Id        int64  `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Token     string `json:"token"`
	}{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     authToken,
	}
	render.Status(req, http.StatusCreated)
	render.JSON(resp, req, respStruct)

	return
}

func (controller *AuthController) Login(resp http.ResponseWriter, req *http.Request) {
	const op = "controllers.auth_controller.Login"

	var loginRequest userRequests.LoginRequest
	if decodeRequest(resp, req, op, &loginRequest) == false {
		return
	}

	if validate(resp, req, &loginRequest) == false {
		return
	}

	token, err := controller.UserService.Login(req.Context(), loginRequest)
	if err != nil {
		render.Status(req, http.StatusUnprocessableEntity)
		render.JSON(resp, req, response.Error(err.Error()))

		return
	}

	render.JSON(resp, req, map[string]string{
		"token": token,
	})
}
