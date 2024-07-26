package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
	"net/http"
	"time"
)

var pacientes = map[string]*models.Paciente{}

func init() {
	pacientes["08741129407"] = &models.Paciente{
		ID:        1,
		Nome:      "Rhuan",
		Sobrenome: "Dantas",
		Cpf:       "08741129407",
		Email:     "rhuan@gmail.com",
		Telefone:  "5581988776655",
		Latitude:  0,
		Longitude: 0,
		UsuarioID: 0,
		CreatedAt: time.Time{},
	}
}

// loginHandler handles the /login endpoint
func GetPaciente(c echo.Context) error {
	cpf := c.Param("cpf")
	paciente, found := pacientes[cpf]
	if !found {
		return c.JSON(http.StatusNotFound, errors.New("paciente not found"))
	}
	// For demonstration, just return the received request
	return c.JSON(http.StatusOK, paciente)
}
