package DI

import (
	"github.com/victorts1991/fiap-pos-tech-hackaton/auth"
	"github.com/victorts1991/fiap-pos-tech-hackaton/db"
	"github.com/victorts1991/fiap-pos-tech-hackaton/handlers"
)

type Dependencies struct {
	Database     *db.Database
	LoginHandler *handlers.Login
	Token        auth.Token
}

func NewDependencies() *Dependencies {
	database := db.NewDatabase()
	token := auth.NewJwtToken()
	login := handlers.NewLogin(database.Usuarios, token)
	return &Dependencies{
		LoginHandler: login,
		Database:     database,
		Token:        token,
	}
}
