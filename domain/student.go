package domain

type Student struct {
	Name     string `json:"name" bson:"name"`
	SchoolID string `json:"id" bson:"id"`
}
