package main

import (
	"fiber-mongo-api/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init() {
	fmt.Println("Go routine main.go")
}

func main() {
	app := fiber.New()
	// configs.ConnectDB()
	routes.UserRoute(app)
	routes.NotesRoutes(app)
	app.Listen(":8080")
}

// configs is used for modularizing project configration files
// controllers is used for modularizing application logics
// models is used for modularizing data and database logics
// responses is for modularizing files describing the response we want our API to give
// routes is for modularizing URL pattern and handler information.
