package controller

import (
	"go-fiber-api/data/request"
	"go-fiber-api/data/response"
	"go-fiber-api/helper"
	"go-fiber-api/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	noteService service.NoteService
}

func NewNoteController(service service.NoteService) *NoteController {
	return &NoteController{
		noteService: service,
	}
}

// CreateNote
func (controller *NoteController) Create(ctx *fiber.Ctx) error {
	req := request.CreateNoteRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)

	controller.noteService.Create(req)

	res := response.Response{
		Code:    201,
		Status:  "OK",
		Message: "Note created",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(res)
}

// UpdateNote
func (controller *NoteController) Update(ctx *fiber.Ctx) error {
	req := request.UpdateNoteRequest{}
	err := ctx.BodyParser(&req)
	helper.ErrorPanic(err)

	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	req.Id = id

	controller.noteService.Update(req)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Note updated",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

// DeleteNote
func (controller *NoteController) Delete(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	controller.noteService.Delete(id)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Note Deleted",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

// FindById
func (controller *NoteController) FindById(ctx *fiber.Ctx) error {
	noteId := ctx.Params("noteId")
	id, err := strconv.Atoi(noteId)
	helper.ErrorPanic(err)

	note := controller.noteService.FindById(id)

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Note Found",
		Data:    note,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

// FindAll
func (controller *NoteController) FindAll(ctx *fiber.Ctx) error {
	notes := controller.noteService.FindAll()

	res := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Notes Found",
		Data:    notes,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
