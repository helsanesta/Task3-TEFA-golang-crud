package Model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Material struct {
	ID          primitive.ObjectID
	CourseName   string
	MaterialName string
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Course struct {
	ID          primitive.ObjectID
	CourseName  string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
