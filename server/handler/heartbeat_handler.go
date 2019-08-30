package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// HandleHeartbeat is healthcheck handler
func HandleHeartbeat(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
