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

type NotesModel struct {
	Title    string             `bson:"title" json:"title" validate:"required"`
    Content  string             `bson:"content" json:"content" validate:"required"`
    Priority int                `bson:"priority" json:"priority" validate:"required"`
    Tags     []string           `bson:"tags,omitempty" json:"tags,omitempty"`
}

func init() {
	fmt.Println("go routinen notes model")
	var notesCollection *mongo.Collection = configs.GetCollection(configs.DB, "notes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	TitleIndex := mongo.IndexModel{
		Keys: bson.M{
			"title": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	ContentIndex := mongo.IndexModel{
		Keys: bson.M{
			"content": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	//This is for single field index
	// if _, err := notesCollection.Indexes().CreateOne(ctx, titleIndex); err != nil {
	// 	log.Fatal(err)
	// }
	if _, err := notesCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{TitleIndex, ContentIndex}); err != nil {
		log.Fatal(err)
	}
}	
