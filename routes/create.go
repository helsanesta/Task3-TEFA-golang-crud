package routes

import (
	"context"
	getcollection "go-crud/Collection"
	database "go-crud/databases"
	model "go-crud/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCourse(c *gin.Context) {
	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	course := new(model.Course)
	defer cancel()

	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	coursePayload := model.Course{
		ID:          primitive.NewObjectID(),
		CourseName:  course.CourseName,
		Description: course.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := courseCollection.InsertOne(ctx, coursePayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
