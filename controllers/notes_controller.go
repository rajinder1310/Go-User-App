package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var notesCollection *mongo.Collection = configs.GetCollection(configs.DB, "notes")
var notes models.NotesModel

// var validate = validator.New()

func CreateNotes(c *fiber.Ctx) error {
	fmt.Println("Enter in notes collection")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.BodyParser(&notes); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	if validationError := validate.Struct(&notes); validationError != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationError.Error()}})
	}

	newNotes := models.NotesModel{
		Title:    notes.Title,
		Content:  notes.Content,
		Priority: notes.Priority,
		Tags:     notes.Tags,
	}

	result, err := notesCollection.InsertOne(ctx, newNotes)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return c.Status(fiber.StatusConflict).JSON(responses.UserResponse{
				Status:  fiber.StatusConflict,
				Message: "error",
				Data:    &fiber.Map{"data": "Duplicate key error"},
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
