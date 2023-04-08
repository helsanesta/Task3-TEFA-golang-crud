package routes

import (
	"context"
	getcollection "go-crud/Collection"
	database "go-crud/databases"
	model "go-crud/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateCourse(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")

	courseID := c.Param("courseID")
	var course model.Course

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(courseID)

	if err := c.BindJSON(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"coursename": course.CourseName, "description": course.Description}
	edited["updatedAt"] = time.Now()

	result, err := courseCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
