package main

import (
	"fmt"
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

	fmt.Println("getting here 1")

	if err := c.Bind(&user); err != nil {
		fmt.Println("getting here 2")
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println("getting here 3")

	if user.Username != "c137@onecause.com" || user.Password != "#th@nH@rm#y#r!$100%D0p#" || user.Token != timeStamp {
		fmt.Println("getting here 4")
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	fmt.Println("getting here 5")

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
