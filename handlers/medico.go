package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/victorts1991/fiap-pos-tech-hackaton/models"
)

var medicos = map[string]*models.Medico{
	"123456": &models.Medico{
		ID:            1,
		CRM:           "123456",
		Nome:          "Dr. Fulano",
		Especialidade: "Cardiologista",
		Email:         "drfulano@gmail.com",
		Avaliacao:     4.3,
		Telefone:      "5581988776655",
		Latitude:      41.3212,
		Longitude:     33.0323,
		CreatedAt:     time.Now(),
		UsuarioID:     2,
		Horarios: []models.Horario{
			models.Horario{
				ID:       1,
				MedicoID: 1,
				Data:     "2024-07-02T15:00",
				Status:   "Disponível",
			},
			models.Horario{
				ID:       2,
				MedicoID: 1,
				Data:     "2024-07-02T16:00",
				Status:   "Disponível",
			},
			models.Horario{
				ID:       3,
				MedicoID: 1,
				Data:     "2024-07-02T17:00",
				Status:   "Disponível",
			},
			models.Horario{
				ID:       4,
				MedicoID: 1,
				Data:     "2024-07-02T18:00",
				Status:   "Disponível",
			},
		},
	},
}

// loginHandler handles the /login endpoint
func GetMedicos(c echo.Context) error {
	var result []*models.Medico
	for _, m := range medicos {
		result = append(result, m)
	}
	return c.JSON(http.StatusOK, result)

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
	horarios[h.ID] = &h
	idCounter++
	return c.JSON(http.StatusCreated, h)
}

func GetHorarioByMedicoID(c echo.Context) error {
	medicoID, _ := strconv.Atoi(c.Param("medico_id"))
	var result []*models.Horario
	for _, h := range horarios {
		if h.MedicoID == medicoID {
			result = append(result, h)
		}
	}
	return c.JSON(http.StatusOK, result)
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
