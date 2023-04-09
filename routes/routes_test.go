package routes

import (
	"context"
	"fmt"
	"testing"
	"time"

	// "github.com/cornelk/hashmap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	getcollection "go-crud/Collection"
	database "go-crud/databases"
	model "go-crud/model"
)

func BenchmarkCRUD(b *testing.B) {
	// Set up database connection
	var DB = database.ConnectDB()
	var courseCollection = getcollection.GetCollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, numRecords := range []int{100000} {
		b.Run(fmt.Sprintf("numRecords=%d", numRecords), func(b *testing.B) {
			// Create sample data
			var courses = make(map[primitive.ObjectID]model.Course)
			for i := 0; i < numRecords; i++ {
				course := model.Course{
					ID:          primitive.NewObjectID(),
					CourseName:  fmt.Sprintf("Course %d", i),
					Description: fmt.Sprintf("Description for Course %d", i),
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				courses[course.ID] = course
			}

			// Insert sample data
			var courseDocs []interface{}
			for _, course := range courses {
				courseDocs = append(courseDocs, course)
			}
			_, err := courseCollection.InsertMany(ctx, courseDocs)
			if err != nil {
				b.Fatal(err)
			}

			// Query sample data
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cursor, err := courseCollection.Find(ctx, bson.M{})
				if err != nil {
					b.Fatal(err)
				}
				defer cursor.Close(ctx)

				var results []model.Course
				if err = cursor.All(ctx, &results); err != nil {
					b.Fatal(err)
				}
			}

			// Clean up sample data
			_, err = courseCollection.DeleteMany(ctx, bson.M{})
			if err != nil {
				b.Fatal(err)
			}
		})
	}
}

// package routes

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"

// 	getcollection "go-crud/Collection"
// 	database "go-crud/databases"
// 	model "go-crud/model"
// )

// func BenchmarkCRUD(b *testing.B) {
// 	// Set up database connection
// 	var DB = database.ConnectDB()
// 	var courseCollection = getcollection.GetCollection(DB, "Posts")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	for _, numRecords := range []int{1000} {
// 		b.Run(fmt.Sprintf("numRecords=%d", numRecords), func(b *testing.B) {
// 			// Create sample data
// 			var courses []model.Course
// 			for i := 0; i < numRecords; i++ {
// 				course := model.Course{
// 					ID:          primitive.NewObjectID(),
// 					CourseName:  fmt.Sprintf("Course %d", i),
// 					Description: fmt.Sprintf("Description for Course %d", i),
// 					CreatedAt:   time.Now(),
// 					UpdatedAt:   time.Now(),
// 				}
// 				courses = append(courses, course)
// 			}

// 			// Insert sample data
// 			for _, course := range courses {
// 				_, err := courseCollection.InsertOne(ctx, course)
// 				if err != nil {
// 					b.Fatal(err)
// 				}
// 			}

// 			// Query sample data
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				cursor, err := courseCollection.Find(ctx, bson.M{})
// 				if err != nil {
// 					b.Fatal(err)
// 				}
// 				defer cursor.Close(ctx)

// 				var results []model.Course
// 				if err = cursor.All(ctx, &results); err != nil {
// 					b.Fatal(err)
// 				}
// 			}

// 			// Clean up sample data
// 			_, err := courseCollection.DeleteMany(ctx, bson.M{})
// 			if err != nil {
// 				b.Fatal(err)
// 			}
// 		})
// 	}
// }
