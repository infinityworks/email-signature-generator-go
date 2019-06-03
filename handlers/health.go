package handlers

import (
	"flag"
	"net/http"

	"github.com/infinityworks/email-signature-generator/models"
	"github.com/labstack/echo"
)

// Health handler for reporting app status and version
func Health(context echo.Context) error {
	version := flag.Lookup("version").Value.(flag.Getter).Get().(string)
	health := models.NewHealth("OK", version)
	return context.JSON(http.StatusOK, &health)
}
