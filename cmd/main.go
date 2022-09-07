package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    int    `json:"token"`
}

func handleUserValidation(c echo.Context) error {
	var user User
	timeStamp := time.Now().Hour()*100 + time.Now().Minute()

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	if user.Username != "c137@onecause.com" || user.Password != "#th@nH@rm#y#r!$100%D0p#" || user.Token != timeStamp {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusCreated, "Login Success")
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handleUserValidation)

	e.Logger.Fatal(e.Start(":3030"))
}
