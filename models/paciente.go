package models

import (
	"bytes"
	"errors"
	"time"
	"unicode"

	"github.com/paemuri/brdoc"
)

// Paciente struct
type Paciente struct {
	ID        int       `json:"id"`
	Nome      string    `json:"nome" validate:"required"`
	Sobrenome string    `json:"sobrenome" validate:"required"`
	Cpf       string    `json:"cpf" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Telefone  string    `json:"telefone" validate:"required"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	UsuarioID int       `json:"usuario_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Paciente) ValidateCPF() error {
	if !brdoc.IsCPF(p.Cpf) {
		return errors.New("cpf invalido")
	}

	p.limpaCaracteresEspeciais()

	return nil
}

func (p *Paciente) limpaCaracteresEspeciais() {
	buf := bytes.NewBufferString("")
	for _, r := range p.Cpf {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	p.Cpf = buf.String()
}
