package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `json:"mongo_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	SchoolID string             `json:"id" bson:"id"`
}
