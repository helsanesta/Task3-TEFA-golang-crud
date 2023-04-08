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
)

func GetAllCourses(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")

	cursor, err := courseCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	defer cursor.Close(ctx)

	var courses []model.Course
	if err = cursor.All(ctx, &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": courses})
}
