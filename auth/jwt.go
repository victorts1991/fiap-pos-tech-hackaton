package auth

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/victorts1991/fiap-pos-tech-hackaton/errors"
)

//go:generate mockgen -source=$GOFILE -package=mock_auth -destination=../../../../test/mock/auth/$GOFILE

type Token interface {
	GenerateToken(tipo string, payload any) (string, error)
	VerifyToken(next echo.HandlerFunc) echo.HandlerFunc
	PermissaoPaciente(next echo.HandlerFunc) echo.HandlerFunc
	PermissaoMedico(next echo.HandlerFunc) echo.HandlerFunc
	PermissaoTodosUsuarios(next echo.HandlerFunc) echo.HandlerFunc
}

type JwtToken struct{}

func NewJwtToken() Token {
	return &JwtToken{}
}

type jwtCustomClaims struct {
	Usuario any    `json:"payload"`
	Tipo    string `json:"tipo"`
	jwt.MapClaims
}

func (jt *JwtToken) GenerateToken(tipo string, payload any) (string, error) {
	secret := []byte(os.Getenv("AUTH_SECRET"))
	claims := jwt.MapClaims{
		"payload": payload,
		"tipo":    tipo,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

func (jt *JwtToken) VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := jt.getToken(c)
		if tokenStr == "" {
			return errors.HandleError(c, errors.Unauthorized.New("authentication key not found"))
		}
		claims := jwt.MapClaims{}
		secret := []byte(os.Getenv("AUTH_SECRET"))
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			return errors.Unauthorized.New(err.Error())
		}

		if !tkn.Valid {
			return errors.Unauthorized.New("authentication is not valid")
		}

		permissao := c.Get("permissao")
		if permissao != "todos" && permissao != claims["tipo"] {
			return errors.Unauthorized.New("tipo não permitido para esta chamada")
		}

		c.Set("usuario", claims["payload"])

		return next(c)
	}
}

func (jt *JwtToken) PermissaoPaciente(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("permissao", "paciente")
		return next(c)
	}
}

func (jt *JwtToken) PermissaoMedico(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("permissao", "medico")
		return next(c)
	}
}

func (jt *JwtToken) PermissaoTodosUsuarios(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("permissao", "todos")
		return next(c)
	}
}

func (jt *JwtToken) getToken(c echo.Context) string {
	tokenStr := ""
	if bearer := c.Request().Header.Get("Authorization"); bearer != "" {
		if strings.Contains(bearer, "Bearer") {
			tokenStr = strings.Split(bearer, " ")[1]
		}
	}

	if tokenStr == "" {
		tokenStr = c.Request().Header.Get("token")
	}

	return tokenStr
}
