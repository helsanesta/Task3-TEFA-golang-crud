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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMaterial(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var DB = database.ConnectDB()
	var materialCollection = getcollection.GetCollection(DB, "Materials")

	cursor, err := materialCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	defer cursor.Close(ctx)

	var materials []model.Material
	if err = cursor.All(ctx, &materials); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "Data": materials})
}

func CreateMaterial(c *gin.Context) {
	var DB = database.ConnectDB()
	var materialCollection = getcollection.GetCollection(DB, "Materials")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	material := new(model.Material)
	defer cancel()

	if err := c.BindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	materialPayload := model.Material{
		ID:           primitive.NewObjectID(),
		CourseName:   material.CourseName,
		MaterialName: material.MaterialName,
		Description:  material.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := materialCollection.InsertOne(ctx, materialPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Material upload successfully", "Data": map[string]interface{}{"data": result}})
}

func DeleteMaterial(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	materialID := c.Param("materialID")

	var materialCollection = getcollection.GetCollection(DB, "Materials")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(materialID)
	result, err := materialCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Material deleted successfully", "Data": res})
}

func UpdateMaterial(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var materialCollection = getcollection.GetCollection(DB, "Materials")

	materialID := c.Param("materialID")
	var material model.Material

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(materialID)

	if err := c.BindJSON(&material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"CourseName": material.CourseName, "MaterialName": material.MaterialName, "description": material.Description}
	edited["updatedAt"] = time.Now()

	result, err := materialCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

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
