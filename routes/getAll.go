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

	var courses []map[string]interface{}
	for cursor.Next(ctx) {
		var course model.Course
		err := cursor.Decode(&course)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		// Store course data in a map
		data := make(map[string]interface{})
		data["id"] = course.ID.Hex()
		data["coursename"] = course.CourseName
		data["description"] = course.Description
		data["created_at"] = course.CreatedAt
		data["updated_at"] = course.UpdatedAt

		courses = append(courses, data)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": courses})
}

// WITHOUT MAP / HASH TABLE

// package routes

// import (
// 	"context"
// 	getcollection "go-crud/Collection"
// 	database "go-crud/databases"
// 	model "go-crud/model"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func GetAllCourses(c *gin.Context) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var DB = database.ConnectDB()
// 	var courseCollection = getcollection.GetCollection(DB, "Posts")

// 	cursor, err := courseCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var courses []model.Course
// 	if err = cursor.All(ctx, &courses); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": courses})
// }
