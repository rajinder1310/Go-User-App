package models

import (
	"context"
	"fiber-mongo-api/configs"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserModel struct {
	Name     string `bson:"name" json:"name" validate:"required"`
	Location string `bson:"location,omitempty"`
	Title    string `bson:"title" json:"title" validate:"required"`
}

func init() {
	fmt.Println("Enter in go routine : user.module.go")
	var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Create index on email field
	compositeIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1}, 
			{Key: "location", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}
	if _, err := userCollection.Indexes().CreateOne(ctx, compositeIndex); err != nil {
		fmt.Println("UnConfirmed", err)
		log.Fatal(err)
	}

}
