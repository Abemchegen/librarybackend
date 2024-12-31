package domain

type LendBookRequest struct {
	BookID      string `json:"book_id" binding:"required"`
	StudentID   string `json:"student_id" binding:"required"`
	StudentName string `json:"student_name" binding:"required"`
	LentDate    string `json:"lent_date" binding:"required"`
	DueDate     string `json:"due_date" binding:"required"`
	LentType    string `json:"lent_type" binding:"required"`
}

type ReturnBookRequest struct {
	BookID          string `json:"book_id" binding:"required"`
	StudentID       string `json:"student_id" binding:"required"`
	ReturnStatus    string `json:"return_status" binding:"required"`
	ReturnCondition string `json:"return_condition" binding:"required"`
}
