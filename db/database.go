package db

import "github.com/victorts1991/fiap-pos-tech-hackaton/models"

type Database struct {
	Usuarios    []*models.Usuario
	Medicos     map[string]*models.Medico
	Horarios    map[int]*models.Horario
	Pacientes   map[string]*models.Paciente
	Prontuarios map[int]*models.Prontuario
}

func NewDatabase() *Database {
	return &Database{
		Usuarios: []*models.Usuario{
			&models.Usuario{
				ID:       1,
				UserName: "paciente1",
				Senha:    "paciente1",
				Tipo:     "paciente",
			},
			&models.Usuario{
				ID:       1,
				UserName: "medico1",
				Senha:    "medico1",
				Tipo:     "medico",
			},
		},
		Medicos:  make(map[string]*models.Medico),
		Horarios: make(map[int]*models.Horario),
		Pacientes: map[string]*models.Paciente{
			"12345678901": &models.Paciente{
				ID:   1,
				Nome: "Paciente 1",
				Cpf:  "12345678901",
			},
		},
		Prontuarios: make(map[int]*models.Prontuario),
	}
}

func seedUsuario() *models.Usuario {
	return &models.Usuario{
		ID:       1,
		UserName: "paciente1",
		Senha:    "paciente1",
		Tipo:     "paciente",
	}
}
