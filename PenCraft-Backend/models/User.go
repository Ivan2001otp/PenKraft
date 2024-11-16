package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	User_id       primitive.ObjectID `json:"user_id" bson:"user_id"`
	First_name    string             `json:"first_name" bson:"first_name" validate:"required"`
	Last_name     string             `json:"last_name" bson:"last_name" validate:"required"`
	Bio           *string            `json:"bio_data,omitempty" bson:"bio_data,omitempty"`
	Email         string             `json:"email" bson:"email" validate:"required,email"`
	Password      string             `json:"password" bson:"password"`
	Profile_image *string            `json:"profile_image,omitempty" bson:"profile_image,omitempty"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
}
