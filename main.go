package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/victorts1991/fiap-pos-tech-hackaton/handlers"
)

//	@title			Hackathon
//	@version		1.0.0
//	@description	This is a documentation of all endpoints in the API.

// @host		localhost:3000
// @BasePath	/
// @schemes http
// @produce json
// @securityDefinitions.apikey	JWT
// @in							header
// @name						token
func main() {
	e := echo.New()

	// Initialize the database
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/login", handlers.Login)
	e.GET("/pacientes/:cpf", handlers.GetPaciente)
	e.GET("/medicos", handlers.GetMedicos)
	//horario
	e.POST("/horarios", handlers.CreateHorario)
	e.GET("/horarios/:medico_id", handlers.GetHorario)
	e.PUT("/horarios/:id", handlers.UpdateHorario)
	e.DELETE("/horarios/:id", handlers.DeleteHorario)
	//prontuario
	e.POST("/prontuarios", handlers.CreateProntuario)
	e.GET("/prontuarios/:id", handlers.GetProntuario)
	e.PUT("/prontuarios/:id", handlers.UpdateProntuario)
	e.DELETE("/prontuarios/:id", handlers.DeleteProntuario)
	//consulta
	e.POST("/consultas", handlers.CreateConsulta)
	e.GET("/consultas/:id", handlers.GetConsulta)
	e.PUT("/consultas/:id", handlers.UpdateConsulta)
	e.DELETE("/consultas/:id", handlers.DeleteConsulta)
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}

// CustomValidator wraps the go-playground validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate implements the echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
