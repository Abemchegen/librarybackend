package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	ID          primitive.ObjectID `json:"mongo_id" bson:"_id,omitempty"`
	StudentName string             `json:"studentname" bson:"studentname"`
	StudentID   string             `json:"studentid" bson:"studentid"`
	EntryTime   time.Time          `json:"entrytime" bson:"entrytime"`
	LeaveTime   *time.Time         `json:"leavetime" bson:"leavetime"`
}
