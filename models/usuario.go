package models

import "time"

// Usuario struct
type Usuario struct {
	ID              int       `json:"id"`
	Nome            string    `json:"nome" validate:"required"`
	UserName        string    `json:"user_name" validate:"required"`
	Email           string    `json:"email" validate:"required"`
	Senha           string    `json:"senha" validate:"required"`
	Tipo            string    `json:"tipo" validate:"required"`
	Token           string    `json:"token"`
	TokenExpiration time.Time `json:"token_expiration"`
}

type LoginRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Senha    string `json:"senha" validate:"required"`
	Tipo     string `json:"tipo" validate:"required" oneof:"medico paciente"`
}
