package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

type (
	user struct {
		Username int `json:"username"`
		Password string `json:"password"`
		Token int `json:"token"`
	}
)

func createUser(c echo.Context) error {
	// e.Logger.
	// u: user = c.
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusCreated, '')
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", createUser)

	e.Logger.Fatal(e.Start(":3030"))
}