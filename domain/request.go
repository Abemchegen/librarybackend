package domain

type LendBookRequest struct {
	BookID      string `json:"bookid" bson:"bookid"`
	StudentID   string `json:"studentid" bson:"studentid"`
	StudentName string `json:"studentname" bson:"studentname"`
	LentDate    string `json:"lentdate" bson:"lentdate"`
	DueDate     string `json:"duedate" bson:"duedate"`
	LentType    string `json:"lenttype" bson:"lenttype"`
}

type ReturnBookRequest struct {
	BookID          string `json:"bookid" bson:"bookid"`
	StudentID       string `json:"studentid" bson:"studentid"`
	ReturnStatus    string `json:"returnstatus" bson:"returnstatus"`
	ReturnCondition string `json:"returncondition" bson:"returncondition"`
	ReturnDate      string `json:"returndate" bson:"returndate"`
}
