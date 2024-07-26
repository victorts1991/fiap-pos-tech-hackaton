package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/victorts1991/fiap-pos-tech-hackaton/DI"
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

	//dependencies
	dependencies := DI.NewDependencies()

	// Routes
	e.GET("/liveness", liveness)
	e.POST("/login", dependencies.LoginHandler.LoginHandler)
	e.GET("/pacientes/:cpf", dependencies.Token.PermissaoPaciente(dependencies.Token.VerifyToken(handlers.GetPaciente)))
	e.GET("/medicos", handlers.GetMedicos)
	//horarios
	e.POST("/horarios", handlers.CreateHorario)
	e.GET("/horarios/:medico_id", handlers.GetHorarioByMedicoID)
	e.PUT("/horarios/:id", handlers.UpdateHorario)
	e.DELETE("/horarios/:id", handlers.DeleteHorario)
	//prontuarios
	e.POST("/prontuarios", handlers.CreateProntuario)
	e.GET("/prontuarios/:paciente_id", handlers.GetProntuarioByPacienteID)
	e.PUT("/prontuarios/:id", handlers.UpdateProntuario)
	e.DELETE("/prontuarios/:id", handlers.DeleteProntuario)
	//consultas
	e.PATCH("/consultas/:id", handlers.AtualizaSolicitacaoConsulta)
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

// Liveness godoc
// @Summary Show the status of http.
// @Description get the status of http.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {string} string "token"
// @Router /liveness [get]
func liveness(c echo.Context) error {
	response := make(map[string]bool)
	response["status"] = true
	return c.JSON(http.StatusOK, response)
}
