package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Home!")
}

func main() {
	e := echo.New()
	e.GET("/", homeHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
