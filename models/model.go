package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	Thread      `json:"thread,omitempty"`
}

type Thread struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Message   string             `json:"message,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
