package models

import "time"

type Prontuario struct {
	ID         int       `json:"id"`
	PacienteID int       `json:"paciente_id" validate:"required"`
	Files      []string  `json:"files"`
	Observacao string    `json:"observacao"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
