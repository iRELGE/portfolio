package apimanager

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"rabie.com/portfolio/models"
)

// ValidateUserData use this methode
func ValidateUserData(user models.User, c echo.Context) bool {
	u := new(UserRegister)
	u.Email = user.Email
	u.Password = user.Password
	if err := c.Bind(u); err != nil {
		return false
	}

	if err := c.Validate(u); err != nil {
		fmt.Println(u, "error2")
		return false

	}
	return true

}
