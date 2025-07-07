package handlers

import (
	"net/http"

	"course-go/modules/user/interfaces"
	"course-go/modules/user/models"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(c echo.Context, userService interfaces.UserService) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	createdUser, err := userService.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createdUser)
}
