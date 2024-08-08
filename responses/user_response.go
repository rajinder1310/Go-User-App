package responses

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func init() {
	fmt.Println("Go routine responses")
}
type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data 	*fiber.Map `json:"data"`
}