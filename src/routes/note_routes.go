package routes

import (
	"github.com/EliasPereyra/go-react-crud/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func NoteRoute(app *fiber.App) {
	app.Get("/notes", controllers.GetNotes)
}
