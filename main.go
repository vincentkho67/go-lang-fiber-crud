package main

import (
	"fmt"
	"go-fiber-api/config"
	"go-fiber-api/controller"
	"go-fiber-api/model"
	"go-fiber-api/repository"
	"go-fiber-api/router"
	"go-fiber-api/service"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Run Service ...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("no env var detected", err)
	}

	// Init DB
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Init Controller
	noteController := controller.NewNoteController(noteService)

	// routes
	routes := router.NewRouter(noteController)

	app := fiber.New()
	app.Mount("/api/v1", routes)

	log.Fatal(app.Listen(":3000"))
}
