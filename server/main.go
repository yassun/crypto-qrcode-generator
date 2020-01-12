package main

import (
	"os"
	"server/handler"
	"server/validator"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := os.Args[1]
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = validator.NewCustomValidator()
	e.GET("/heartbeat", handler.HandleHeartbeat)
	e.POST("/api/v1/generate-qr/btc", handler.GenerateBtcQR)
	e.Logger.Fatal(e.Start(":" + port))
}
