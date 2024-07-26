package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
	"net/http"
	"strconv"
	"time"
)

var consultas = map[int]*models.Consulta{}

func CreateConsulta(c echo.Context) error {
	var con models.Consulta
	if err := c.Bind(&con); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := c.Validate(&con); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	con.ID = idCounter
	con.CreatedAt = time.Now().Format(time.RFC3339)
	consultas[con.ID] = &con
	idCounter++
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
