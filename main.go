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

	/* MATERIALS */
	// called as localhost:3000/materials
	router.GET("/materials", routes.GetAllMaterial)

	// called as localhost:3000/materials/:id
	router.GET("/material/:materialID", routes.GetOneMaterial)

	// called as localhost:3000/materials
	router.POST("/materials", routes.CreateMaterial)

	// called as localhost:3000/materials/:id
	router.PUT("/material/:materialID", routes.UpdateMaterial)

	// called as localhost:3000/materials/:id
	router.DELETE("/material/:materialID", routes.DeleteMaterial)

	router.Run("localhost: 3000")
}
