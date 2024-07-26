package models

import "time"

type Medico struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome" validate:"required"`
	Sobrenome string    `json:"sobrenome" validate:"required"`
	CRM       string    `json:"crm" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Telefone  string    `json:"telefone" validate:"required"`
	Latitude  float64   `json:"latitude"`
	UsuarioID int       `json:"usuario_id" validate:"required"`
	Horarios  []Horario `json:"horarios"`
	Longitude float64   `json:"longitude"`
	CreatedAt string    `json:"created_at"`
}

type Horario struct {
	ID        int       `json:"id"`
	MedicoID  int       `json:"medico_id" validate:"required"`
	Data      time.Time `json:"data" validate:"required"`
	Status    string    `json:"status" validate:"required"`
	CreatedAt string    `json:"created_at"`
}
