package controllers

import (
	"auth-service/helpers"
	"auth-service/web/models"
	"auth-service/web/responses"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserController interface {
	Login(w http.ResponseWriter, r *http.Request)
	Ping(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	Validate *validator.Validate
}

func NewUserController(validate *validator.Validate) UserController {
	return &userControllerImpl{
		Validate: validate,
	}
}

func (c *userControllerImpl) Ping(w http.ResponseWriter, r *http.Request) {
	webResponse := responses.WebResponse{
		Code:   http.StatusOK,
		Status: "Hello there",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *userControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userRequestLogin := models.UserLogin{}
	helpers.ToRequestBody(r, &userRequestLogin)

	err := c.Validate.Struct(userRequestLogin)
	helpers.PanicIfError(err)

	if userRequestLogin.Email == "admin@admin.com" && userRequestLogin.Password == "asdasd" {
		webResponse := responses.WebResponse{
			Code:   http.StatusOK,
			Status: "Login Success",
			Data:   userRequestLogin,
		}
		helpers.WriteResponseBody(w, webResponse)
	} else {
		webResponse := responses.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid Credential",
		}
		helpers.WriteResponseBody(w, webResponse)
	}
}
