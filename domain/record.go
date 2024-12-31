package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Student         Student            `json:"student" bson:"student"`
	Book            Book               `json:"Book" bson:"Book"`
	LentDate        string             `json:"lent_at" bson:"lent_at"`
	ReturnDue       string             `json:"returned_at" bson:"returned_at"`
	LentType        string             `json:"lent_type" bson:"lent_type"`
	ReturnDate      string             `json:"return_date" bson:"return_date"`
	ReturnStatus    string             `json:"return_status" bson:"return_status"`
	ReturnCondition string             `json:"return_condition" bson:"return_condition"`
}

// type RecordUseCase interface {
// 	CreateRecord(student Student, book Book, lentdate string, duedate string, lenttype string, returnstatus string, returncondition string) (Record, error)
// 	GetAllRecord() ([]Record, error)
// 	UpdateRecord(studentid string, bookid string, returnstatus string, returncondition string) (Record, error)
// }

type RecordRepository interface {
	CreateRecord(student Student, book Book, lentdate string, duedate string, lenttype string, returnstatus string, returncondition string) (Record, error)
	GetAllRecord() ([]Record, error)
	UpdateRecord(studentid string, bookid string, returnstatus string, returncondition string) (Record, error)
}
