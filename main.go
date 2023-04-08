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

	/* MATERIALS */
	// called as localhost:3000/materials
	router.GET("/materials", routes.GetAllMaterial)

	// called as localhost:3000/materials
	router.POST("/materials", routes.CreateMaterial)

	// called as localhost:3000/materials/:id
	router.PUT("/materials/:materialID", routes.UpdateMaterial)

	// called as localhost:3000/materials/:id
	router.DELETE("/materials/:materialID", routes.DeleteMaterial)

	router.Run("localhost: 3000")
}
