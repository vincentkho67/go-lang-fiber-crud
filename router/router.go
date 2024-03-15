package router

import (
	"go-fiber-api/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Success",
			"message": "Welcome to Enigma Camp",
		})
	})

	router.Route("/notes", func(router fiber.Router) {
		router.Post("", noteController.Create)
		router.Get("", noteController.FindAll)
	})

	router.Route("/notes/:noteId", func(router fiber.Router) {
		router.Delete("", noteController.Delete)
		router.Get("", noteController.FindById)
		router.Put("", noteController.Update)
	})

	return router
}
