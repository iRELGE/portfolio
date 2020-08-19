package apimanager

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"rabie.com/portfolio/createjwt"
	"rabie.com/portfolio/dbmanager"
	"rabie.com/portfolio/img"
	"rabie.com/portfolio/models"
)

// Createproject function to handel project request
func Createproject(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*createjwt.JwtCustomClaims)

	defer c.Request().Body.Close()
	Title := c.FormValue("Title")
	Description := c.FormValue("Description")
	Link := c.FormValue("Link")
	file, err := c.FormFile("Photo")
	if err != nil {
		return c.String(http.StatusBadRequest, "please resend an other photo")
	}
	fileName, err := img.Uploadimg(file, "projects")
	if err != nil {
		return c.String(http.StatusBadRequest, "please resend an other photo")
	}
	newproject := new(models.Project)
	newproject.Title = Title
	newproject.Description = Description
	newproject.Link = Link
	newproject.Photo = fileName
	newproject.UserID = claims.ID
	err = dbmanager.Createproject(*newproject)
	if err != nil {
		return c.String(http.StatusBadRequest, "please try again somthing wrong")
	}
	return c.String(http.StatusOK, "your project has been registred")

}
