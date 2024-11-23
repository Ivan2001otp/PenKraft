package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Blog struct {
	ID         primitive.ObjectID `bson:"_id"`
	Blog_id    string             `json:"blog_id" bson:"blog_id"`
	Title      string             `json:"title" bson:"title" validate:"required"`
	Excerpt    string             `json:"excerpt" bson:"excerpt" validate:"required"`
	Tag_id     string             `json:"tag_id" bson:"tag_id"`
	Is_delete  bool               `json:"is_delete" bson:"is_delete"`
	User_id    string             `json:"user_id" bson:"user_id"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Body       string             `json:"body" bson:"body" validate:"required"`
	Image      string             `json:"image" bson:"image" validate:"required"`
	Slug       string             `json:"slug" bson:"slug"`
}
