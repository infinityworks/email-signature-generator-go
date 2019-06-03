package handlers

import (
	"net/http"

	"github.com/infinityworks/email-signature-generator/models"
	"github.com/labstack/echo"
)

// GetForm handler for providing the user with a form
func GetForm(c echo.Context) error {
	return c.Render(http.StatusOK, "form.html", map[string]interface{}{})
}

// PostForm handler for taking in the details and sending the actual signature
func PostForm(c echo.Context) error {
	sig := new(models.SignatureDetails)
	if err := c.Bind(sig); err != nil {
		return err
	}
	if err := c.Validate(sig); err != nil {
		return err
		// TODO: Give back useful error message
	}

	return c.Render(http.StatusOK, "signature.html", map[string]interface{}{
		"Signature": sig,
	})
}
