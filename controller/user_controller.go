package controller

import (
	"clothApp-backend/model"
	"clothApp-backend/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
}

type userController struct {
	uu usecases.IUserUsecase
}

func NewUserController(uu usecases.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, userRes)
	}
}
