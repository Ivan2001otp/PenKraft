package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag struct {
	Tag_id primitive.ObjectID `json:"tag_id" bson:"tag_id"`
	Tag_name string `json:"tag_name" bson:"tag_name" validator:"min=1,required"`
	Created_at time.Time `json:"created_at" bson:"created_at"`
	Updated_at time.Time `json:"updated_at" bson:"updated_at"`
}