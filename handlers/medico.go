package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
	"net/http"
	"strconv"
	"time"
)

var medicos = map[string]*models.Medico{}

// loginHandler handles the /login endpoint
func GetMedicos(c echo.Context) error {
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

var horarios = map[int]*models.Horario{}
var idCounter = 1

// Handlers

func CreateHorario(c echo.Context) error {
	var h models.Horario
	if err := c.Bind(&h); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&h); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	h.ID = idCounter
	h.CreatedAt = time.Now().Format(time.RFC3339)
	horarios[h.ID] = &h
	idCounter++
	return c.JSON(http.StatusCreated, h)
}

func GetHorario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if h, ok := horarios[id]; ok {
		return c.JSON(http.StatusOK, h)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Horario not found"})
}

func UpdateHorario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if h, ok := horarios[id]; ok {
		if err := c.Bind(h); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(h); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		horarios[id] = h
		return c.JSON(http.StatusOK, h)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Horario not found"})
}

func DeleteHorario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := horarios[id]; ok {
		delete(horarios, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Horario not found"})
}
