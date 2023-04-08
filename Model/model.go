package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID          primitive.ObjectID
	CourseName  string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
