package handler

import (
	"net/http"
	"todo/app/models"

	"github.com/labstack/echo/v4"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := models.User{}
	user.Email = email
	user.Password = password

	user, err := models.GetUserByEmailAndPassword(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func LoginCheck(c echo.Context) error {
	token := c.FormValue("token")

	user, _ := models.GetUserById(1)
	if user.UUID == token {
		return c.JSON(http.StatusOK, nil)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "but request"})
	}
}
