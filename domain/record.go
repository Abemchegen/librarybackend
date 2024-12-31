package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Student         Student            `json:"student" bson:"student"`
	Book            Book               `json:"book" bson:"book"`
	LentDate        string             `json:"lentat" bson:"lentat"`
	ReturnDue       string             `json:"returnedat" bson:"returnedat"`
	LentType        string             `json:"lenttype" bson:"lenttype"`
	ReturnDate      string             `json:"returndate" bson:"returndate"`
	ReturnStatus    string             `json:"returnstatus" bson:"returnstatus"`
	ReturnCondition string             `json:"returncondition" bson:"returncondition"`
}

type RecordRepository interface {
	CreateRecord(student Student, book Book, lentdate string, duedate string, lenttype string, returnstatus string, returncondition string) (Record, error)
	GetAllRecord() ([]Record, error)
	UpdateRecord(studentid string, bookid string, returndate string, returnstatus string, returncondition string) (Record, error)
	GetRecordByID(studentid string, bookid string) (Record, error)
}
