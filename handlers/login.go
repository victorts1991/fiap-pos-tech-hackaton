package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/victorts1991/fiap-pos-tech-hackaton/auth"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
)

type Login struct {
	usuarios []*models.Usuario
	token    auth.Token
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewLogin(usuarios []*models.Usuario, token auth.Token) *Login {
	return &Login{
		usuarios: usuarios,
		token:    token,
	}
}

// loginHandler handles the /login endpoint
func (l *Login) LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	// For demonstration, just return the received request

	validateTipoRequest := req.Tipo == "medico" || req.Tipo == "paciente"
	if !validateTipoRequest {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid tipo"})
	}

	// Check if user exists
	for _, u := range l.usuarios {
		if u.UserName == req.UserName && u.Senha == req.Senha {
			token, err := l.token.GenerateToken(req.Tipo, u)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})

			}

			return c.JSON(http.StatusOK, LoginResponse{Token: token})
		}
	}

	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid credentials"})
}
