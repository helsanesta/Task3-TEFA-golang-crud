package main

import (
	routes "go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/courses", routes.CreateCourse)

	// called as localhost:3000/getAll
	router.GET("/courses", routes.GetAllCourses)

	// called as localhost:3000/get/{id}
	router.GET("/course/:courseID", routes.ReadOneCourse)

	// called as localhost:3000/update/{id}
	router.PUT("/course/:courseID", routes.UpdateCourse)

	// called as localhost:3000/delete/{id}
	router.DELETE("/course/:courseID", routes.DeleteCourse)

	router.Run("localhost: 3000")
}
