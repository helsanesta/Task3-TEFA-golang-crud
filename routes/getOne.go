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

func ReadOneCourse(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")

	courseID := c.Param("courseID")
	var result model.Course

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(courseID)

	err := courseCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
