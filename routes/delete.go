package routes

import (
	"context"
	getcollection "go-crud/Collection"
	database "go-crud/databases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteCourse(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	courseID := c.Param("courseID")

	var courseCollection = getcollection.GetCollection(DB, "Posts")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(courseID)
	result, err := courseCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Course deleted successfully", "Data": res})
}
