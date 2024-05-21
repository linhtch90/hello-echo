package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	message := ResponseMessage{
		Message:   "Hello echo framework",
		Character: "Lulu",
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":15000"))
}

type ResponseMessage struct {
	Message   string `json:"message"`
	Character string `json:"character"`
}
