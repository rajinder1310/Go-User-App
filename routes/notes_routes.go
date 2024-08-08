package routes

import (
	"fiber-mongo-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init() {
	fmt.Println("Go routine notes route")
}

func NotesRoutes(app *fiber.App) {
	app.Post("/createNotes", controllers.CreateNotes)
}
