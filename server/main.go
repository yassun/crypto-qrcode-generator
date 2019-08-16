package main

import (
	"server/handler"
	"server/validator"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	e.POST("/generate-qr/btc", handler.GenerateBtcQR)
	e.Logger.Fatal(e.Start(":8080"))
}
