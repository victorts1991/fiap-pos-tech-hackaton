package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
	"net/http"
	"strconv"
	"time"
)

var prontuarios = map[int]*models.Prontuario{}

// Handlers

func CreateProntuario(c echo.Context) error {
	var p models.Prontuario
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&p); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	p.ID = idCounter
	p.CreatedAt = time.Now()
	p.UpdatedAt = p.CreatedAt
	prontuarios[p.ID] = &p
	idCounter++
	return c.JSON(http.StatusCreated, p)
}

func GetProntuario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if p, ok := prontuarios[id]; ok {
		return c.JSON(http.StatusOK, p)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Prontuario not found"})
}

func UpdateProntuario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if p, ok := prontuarios[id]; ok {
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		if err := c.Validate(p); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}
		p.UpdatedAt = time.Now()
		prontuarios[id] = p
		return c.JSON(http.StatusOK, p)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Prontuario not found"})
}

func DeleteProntuario(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if _, ok := prontuarios[id]; ok {
		delete(prontuarios, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, echo.Map{"error": "Prontuario not found"})
}
