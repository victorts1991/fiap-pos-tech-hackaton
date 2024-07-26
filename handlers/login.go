package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
	"net/http"
)

// loginHandler handles the /login endpoint
func Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	// For demonstration, just return the received request
	return c.JSON(http.StatusOK, req)
}
