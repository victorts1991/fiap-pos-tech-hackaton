package models

import "time"

type Consulta struct {
	ID         int       `json:"id"`
	PacienteID int       `json:"paciente_id" validate:"required"`
	MedicoID   int       `json:"medico_id" validate:"required"`
	Horario    Horario   `json:"horario" validate:"required"`
	Status     string    `json:"status" validate:"required"`
	Link       string    `json:"link"`
	Observacao string    `json:"observacao"`
	CreatedAt  time.Time `json:"created_at"`
}
