package apimanager

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"rabie.com/portfolio/createjwt"
	"rabie.com/portfolio/dbmanager"
	"rabie.com/portfolio/img"
	"rabie.com/portfolio/models"
)

// RegisterUserapi use this func to register for api
func RegisterUserapi(c echo.Context) (err error) {

	u := new(models.User)
	//close body after finish whit request
	defer c.Request().Body.Close()
	//decoding data based on the Content-Type header.
	if err = c.Bind(u); err != nil {
		return c.JSON(http.StatusInternalServerError, "somthing wrong please try again")
	}
	// data validation using validator
	if check := ValidateUserData(*u, c); check != true {
		return c.JSON(http.StatusBadRequest, "incorrect email fomat or or empty password")
	}
	//check if email user exists

	check, us := dbmanager.RegisterUserdb(*u)
	if check == false {
		return c.JSON(http.StatusOK, us)
	}
	return c.JSON(http.StatusBadRequest, "you email aready exist")

}

// Loginapi it receive email and password to generate a token
func Loginapi(c echo.Context) error {
	useremail := c.FormValue("Email")
	userpassword := c.FormValue("Password")
	exist, userinfo := dbmanager.LoginExist(useremail, userpassword)
	defer c.Request().Body.Close()
	// Throws unauthorized error
	if exist != true {
		return c.JSON(http.StatusUnauthorized, "please try again or contact support")
	}

	t, err := createjwt.Createjwttoken(userinfo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something wrong")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})

}

// Userprofileapi get user profile
func Userprofileapi(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*createjwt.JwtCustomClaims)
	defer c.Request().Body.Close()

	up := dbmanager.Userprofiledb(claims.ID)
	return c.JSON(http.StatusOK, up)

}

// Updateuserprofile : update user profile by id profile
func Updateuserprofile(c echo.Context) error {
	id := c.Param("id")
	Adress := c.FormValue("Adress")
	Name := c.FormValue("Name")
	LastName := c.FormValue("LastName")
	Phone := c.FormValue("Phone")
	file, err := c.FormFile("Photo")
	if err != nil {
		return c.String(http.StatusBadRequest, "please resend an other photo")
	}
	fileName, err := img.Uploadimg(file, "users")
	if err != nil {
		return c.String(http.StatusBadRequest, "please resend an other photo")
	}
	i, _ := strconv.Atoi(id)
	userUpdate := new(models.Userprofile)
	userUpdate.ID = i
	userUpdate.Adress = Adress
	userUpdate.LastName = LastName
	userUpdate.Phone = Phone
	userUpdate.Name = Name
	userUpdate.Photo = fileName
	err = dbmanager.Updateuserprofiledb(*userUpdate)
	if err != nil {
		return c.String(http.StatusBadRequest, "somthing wrong")

	}
	return c.JSON(http.StatusOK, "your profile has been updated")

}
