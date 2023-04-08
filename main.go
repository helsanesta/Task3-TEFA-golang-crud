package main

import (
	routes "go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/create", routes.CreateCourse)

	// called as localhost:3000/getAll
	router.GET("/get", routes.GetAllCourses)

	// called as localhost:3000/get/{id}
	router.GET("/get/:courseID", routes.ReadOneCourse)

	// called as localhost:3000/update/{id}
	router.PUT("/update/:courseID", routes.UpdateCourse)

	// called as localhost:3000/delete/{id}
	router.DELETE("/delete/:courseID", routes.DeleteCourse)

	router.Run("localhost: 3000")
}
