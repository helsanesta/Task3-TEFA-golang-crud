package routes

import (
	"context"
	"fmt"
	getcollection "go-crud/Collection"
	database "go-crud/databases"
	"go-crud/model"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddFaker() {
	var DB = database.ConnectDB()
	var materialCollection = getcollection.GetCollection(DB, "Materials")

	for i := 0; i < 10000; i++ {
		material := model.Material{
			ID:           primitive.NewObjectID(),
			CourseName:   faker.Word(),
			MaterialName: faker.Word(),
			Description:  faker.Sentence(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		_, err := materialCollection.InsertOne(context.Background(), material)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done inserting data")
}

func AddFakerCourse() {
	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")

	for i := 0; i < 10000; i++ {
		course := model.Course{
			ID:          primitive.NewObjectID(),
			CourseName:  faker.Word(),
			Description: faker.Sentence(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		_, err := courseCollection.InsertOne(context.Background(), course)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done inserting data")
}
