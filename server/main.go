package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/sample", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			struct {
				Code int    `json:"code"`
				Text string `json:"text"`
			}{
				Code: http.StatusOK,
				Text: http.StatusText(http.StatusOK),
			},
		)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
