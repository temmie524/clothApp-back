package controller

import (
	"clothApp-backend/model"
	"clothApp-backend/usecases"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IItemController interface {
	GetAllItems(c echo.Context) error
	GetItemById(c echo.Context) error
	CreateItem(c echo.Context) error
	UpdateItem(c echo.Context) error
	DeleteItem(c echo.Context) error
}

type itemController struct {
	iu usecases.IItemUsecase
}

func NewItemController(iu usecases.IItemUsecase) IItemController {
	return &itemController{iu}
}

func (ic *itemController) GetAllItems(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	itemRes, err := ic.iu.GetAllItems(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, itemRes)
}

func (ic *itemController) GetItemById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("itemId")
	itemId, _ := strconv.Atoi(id)

	itemRes, err := ic.iu.GetItemById(uint(userId.(float64)), uint(itemId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, itemRes)
}

func (ic *itemController) CreateItem(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	item := model.Item{}
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	item.UserId = uint(userId.(float64))
	itemRes, err := ic.iu.CreateItem(item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, itemRes)
}

func (ic *itemController) UpdateItem(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("itemId")
	itemId, _ := strconv.Atoi(id)

	item := model.Item{}
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	itemRes, err := ic.iu.UpdateItem(item, uint(userId.(float64)), uint(itemId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, itemRes)
}

func (ic *itemController) DeleteItem(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := c.Param("itemId")
	itemId, _ := strconv.Atoi(id)

	if err := ic.iu.DeleteItem(uint(userId.(float64)), uint(itemId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
