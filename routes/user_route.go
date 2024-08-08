package routes

import (
	"fiber-mongo-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init() {
	fmt.Println("Go routine route")
}
func UserRoute(app *fiber.App) {
	app.Post("/createUser", controllers.CreateUser)
	app.Get("/getUsers", controllers.GetUsersDetails)
	app.Get("/getUserById/:id", controllers.GetUserById)
	app.Put("/updateUserById/:id", controllers.EditUserById)
	app.Delete("/deleteUserId/:id", controllers.DeleteUserById)
}
