package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
)

var consultas = map[int]*models.Consulta{}

func AtualizaSolicitacaoConsulta(c echo.Context) error {
	con := models.Consulta{
		ID:         len(consultas) + 1,
		PacienteID: 1,
		MedicoID:   1,
		Horario: models.Horario{
			ID:       1,
			MedicoID: 1,
			Data:     "2024-07-02T15:00",
			Status:   "Confirmado",
		},
		Status:     "Agendada",
		Observacao: "Consulta de rotina",
		Link:       "https://meet.google.com/abc-def-ghi",
		CreatedAt:  time.Now(),
	}
	return c.JSON(http.StatusCreated, con)
}

func CreateConsulta(c echo.Context) error {
	con := models.Consulta{
		ID:         len(consultas) + 1,
		PacienteID: 1,
		MedicoID:   1,
		Horario: models.Horario{
			ID:       1,
			MedicoID: 1,
			Data:     "2024-07-02T15:00",
			Status:   "Solicitado",
		},
		Status:     "Agendada",
		Observacao: "Consulta de rotina",
		Link:       "https://meet.google.com/abc-def-ghi",
		CreatedAt:  time.Now(),
	}
	return c.JSON(http.StatusCreated, con)
}

func GetConsulta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if con, ok := consultas[id]; ok {
		return c.JSON(http.StatusOK, con)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Consulta not found"})
}

func UpdateConsulta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if con, ok := consultas[id]; ok {
		if err := c.Bind(con); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(con); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		consultas[id] = con
		return c.JSON(http.StatusOK, con)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Consulta not found"})
}

func DeleteConsulta(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := consultas[id]; ok {
		delete(consultas, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Consulta not found"})
}
