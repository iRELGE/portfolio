package apimanager

import (
	"sync"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rabie.com/portfolio/createjwt"
)

// Wg sync goroutine
var Wg sync.WaitGroup

// ServerHeader : create my own custom  that return a hander function
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// set some value to header
		c.Response().Header().Set(echo.HeaderServer, "rabie/1.0")
		c.Response().Header().Set("fake header", "rabie/Portfolio")
		//return a handler function
		return next(c)
	}
}

// Allapi : find all api service here
func Allapi() {

	e := echo.New()
	e.Use(ServerHeader)
	e.Debug = true
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//using use to spesify our custome middleware to e
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/register", RegisterUserapi)
	// Login route
	e.POST("/login", Loginapi)
	// users group
	u := e.Group("/user")
	a := e.Group("/admin")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &createjwt.JwtCustomClaims{},
		SigningKey: []byte("secretToken"),
	}
	u.Use(middleware.JWTWithConfig(config))
	a.Use(middleware.JWTWithConfig(config))
	u.GET("/profile", Userprofileapi)
	u.PUT("/updateprofile/:id", Updateuserprofile)
	a.POST("/createproject", Createproject)
	defer Wg.Done()
	e.Start(":1324")
}
