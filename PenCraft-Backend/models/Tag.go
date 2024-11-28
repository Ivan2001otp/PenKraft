package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	ID primitive.ObjectID `bson:"_id"`
	Tag_id string `json:"tag_id" bson:"tag_id"`
	Tag_name string `json:"tag_name" bson:"tag_name" validation:"required"`
	Is_delete bool `json:"is_delete" bson:"is_delete" validation:"required"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Updated_at time.Time `json:"updated_at" bson:"updated_at"`
	Blog_count *int `json:"blog_count" bson:"blog_count"`
}