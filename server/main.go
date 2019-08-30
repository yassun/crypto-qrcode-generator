package main

import (
	"server/handler"
	"server/validator"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = validator.NewCustomValidator()
	e.POST("/api/v1/generate-qr/btc", handler.GenerateBtcQR)
	e.Logger.Fatal(e.Start(":8000"))
}
